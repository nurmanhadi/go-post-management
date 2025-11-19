package service

import (
	"net/http"
	"post-management/internal/cache"
	"post-management/internal/entity"
	"post-management/internal/event/producer"
	"post-management/internal/repository"
	"post-management/pkg"
	"post-management/pkg/api"
	"post-management/pkg/dto"
	"post-management/pkg/response"
	"strconv"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type PostService struct {
	logger            zerolog.Logger
	validator         *validator.Validate
	postRepository    *repository.PostRepository
	likeRepository    *repository.LikeRepository
	commentRepository *repository.CommentRepository
	postCache         *cache.PostCache
	postProducer      *producer.PostProducer
}

func NewPostService(logger zerolog.Logger, validator *validator.Validate, postRepository *repository.PostRepository, likeRepository *repository.LikeRepository, commentRepository *repository.CommentRepository, postCache *cache.PostCache, postProducer *producer.PostProducer) *PostService {
	return &PostService{
		logger:            logger,
		validator:         validator,
		postRepository:    postRepository,
		likeRepository:    likeRepository,
		commentRepository: commentRepository,
		postCache:         postCache,
		postProducer:      postProducer,
	}
}

// post
func (s *PostService) PostCreate(request *dto.PostAddRequest) error {
	if err := s.validator.Struct(request); err != nil {
		s.logger.Warn().Err(err).Msg("failed to validate request")
		return err
	}
	totalUser, err := api.UserCountById(request.UserId)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed count by id to user service")
		return err
	}
	if totalUser < 1 {
		s.logger.Warn().Msg("user not found")
		return response.Except(404, "user not found")
	}
	timestamp := time.Now()
	post := &entity.Post{
		UserId:      request.UserId,
		Description: request.Description,
		CreatedAt:   timestamp,
		UpdatedAt:   timestamp,
	}
	id, err := s.postRepository.Create(post)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed save to database")
		return err
	}
	go func() {
		data := &dto.EventProducer[dto.EventPostCreatedProducer]{
			Event:     pkg.BROKER_ROUTE_POST_CREATED,
			Timestamp: timestamp,
			Data: dto.EventPostCreatedProducer{
				PostId:       id,
				UserId:       post.UserId,
				Description:  post.Description,
				TotalLike:    0,
				TotalComment: 0,
				CreatedAt:    timestamp,
				UpdatedAt:    timestamp,
			},
		}
		err := s.postProducer.PostCreated(data)
		if err != nil {
			s.logger.Error().Err(err).Msg("failed post created to producer")
			return
		}
	}()
	s.logger.Info().Str("user_id", strconv.Itoa(int(request.UserId))).Msg("post create success")
	return nil
}
func (s *PostService) PostUpdate(id string, request *dto.PostUpdateRequest) error {
	if err := s.validator.Struct(request); err != nil {
		s.logger.Warn().Err(err).Msg("failed to validate request")
		return err
	}
	newId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed parse string to int64")
		return err
	}
	post, err := s.postRepository.FindById(newId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.Warn().Err(err).Msg("post not found")
			return response.Except(http.StatusNotFound, "post not found")
		}
		s.logger.Error().Err(err).Msg("failed find by id to database")
		return err
	}
	if err := s.postCache.DeleteById(newId); err != nil {
		s.logger.Error().Err(err).Msg("failed delete by id to cache")
	}
	if err := s.postRepository.Save(post); err != nil {
		s.logger.Error().Err(err).Msg("failed save to database")
		return err
	}
	s.logger.Info().Str("post_id", id).Msg("post create success")
	return nil
}
func (s *PostService) PostGetById(id string) (*dto.PostResponse, error) {
	newId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed parse string to int64")
		return nil, err
	}
	resp, err := s.postCache.GetById(newId)
	if err != nil {
		if err == memcache.ErrCacheMiss {
			post, err := s.postRepository.FindById(newId)
			if err != nil {
				if err == gorm.ErrRecordNotFound {
					s.logger.Warn().Err(err).Msg("post not found")
					return nil, response.Except(http.StatusNotFound, "post not found")
				}
				s.logger.Error().Err(err).Msg("failed find by id to database")
				return nil, err
			}
			user, err := api.UserGetById(post.UserId)
			if err != nil {
				s.logger.Error().Err(err).Msg("failed get by id to user service")
				return nil, err
			}
			userIds := make([]int64, 0, len(post.Comments))
			if len(post.Comments) != 0 {
				for _, x := range post.Comments {
					userIds = append(userIds, x.UserId)
				}
			}
			body := &dto.ApiUserGetBySliceIdBody{
				Ids: userIds,
			}
			users, err := api.UserGetBySliceId(body)
			if err != nil {
				s.logger.Error().Err(err).Msg("failed get by slice id to user service")
				return nil, err
			}
			comments := make([]dto.CommentResponse, 0, len(post.Comments))
			if len(post.Comments) != 0 {
				for _, x := range post.Comments {
					for _, y := range users {
						if x.UserId == y.Id {
							comments = append(comments, dto.CommentResponse{
								Id:          x.Id,
								Description: x.Description,
								User: dto.UserResponse{
									Id:       y.Id,
									Username: y.Username,
									Name: dto.UserNameInfo{
										FirstName: y.Name.FirstName,
										Lastname:  y.Name.Lastname,
									},
									AvatarUrl: y.AvatarUrl,
								},
								CreatedAt: x.CreatedAt,
								UpdatedAt: x.UpdatedAt,
							})
						}
					}
				}
			}
			resp := &dto.PostResponse{
				Id:           post.Id,
				Description:  post.Description,
				TotalLike:    len(post.Likes),
				TotalComment: len(post.Comments),
				User: dto.UserResponse{
					Id:       user.Id,
					Username: user.Username,
					Name: dto.UserNameInfo{
						FirstName: user.Name.FirstName,
						Lastname:  user.Name.Lastname,
					},
					AvatarUrl: user.AvatarUrl,
				},
				Comments:  comments,
				CreatedAt: post.CreatedAt,
				UpdatedAt: post.UpdatedAt,
			}
			if err := s.postCache.SetById(post.Id, resp); err != nil {
				s.logger.Error().Err(err).Msg("failed set by id to cache")
				return nil, err
			}
			s.logger.Info().Str("post_id", id).Msg("post get by id success")
			return resp, nil
		}
		s.logger.Error().Err(err).Msg("failed get by id to cache")
		return nil, err
	}

	s.logger.Info().Str("post_id", id).Msg("post get by id success")
	return resp, nil
}
func (s *PostService) PostDelete(id string) error {
	newId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed parse string to int64")
		return err
	}
	post, err := s.postRepository.FindById(newId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.Warn().Err(err).Msg("post not found")
			return response.Except(http.StatusNotFound, "post not found")
		}
		s.logger.Error().Err(err).Msg("failed find by id to database")
		return err
	}
	if err := s.postCache.DeleteById(newId); err != nil {
		s.logger.Error().Err(err).Msg("failed delete by id to cache")
	}
	if err := s.postRepository.Delete(post.Id); err != nil {
		s.logger.Error().Err(err).Msg("failed delete post to database")
		return err
	}
	s.logger.Info().Str("post_id", id).Msg("post delete success")
	return nil
}

// like
func (s *PostService) PostLike(request *dto.LikeAddRequest) error {
	if err := s.validator.Struct(request); err != nil {
		s.logger.Warn().Err(err).Msg("failed to validate request")
		return err
	}
	totalPost, err := s.postRepository.CountById(request.PostId)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed count by id to database")
		return err
	}
	if totalPost < 1 {
		s.logger.Warn().Msg("post not found")
		return response.Except(404, "post not found")
	}
	totalUser, err := api.UserCountById(request.UserId)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed count by id to user service")
		return err
	}
	if totalUser < 1 {
		s.logger.Warn().Msg("user not found")
		return response.Except(404, "user not found")
	}
	totalLike, err := s.likeRepository.CountByPostIdAndUserId(request.PostId, request.UserId)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed count by post_id and user_id to database")
		return err
	}
	if totalLike > 0 {
		s.logger.Warn().Msg("like already exists")
		return response.Except(409, "like already exists")
	}
	like := &entity.Like{
		PostId: request.PostId,
		UserId: request.UserId,
	}
	if err := s.postCache.DeleteById(request.PostId); err != nil {
		s.logger.Error().Err(err).Msg("failed delete by id to cache")
	}
	if err := s.likeRepository.Save(like); err != nil {
		s.logger.Error().Err(err).Msg("failed save to database")
		return err
	}
	s.logger.Info().Str("post_id", strconv.Itoa(int(request.PostId))).Msg("post like success")
	return nil
}
func (s *PostService) PostUnlike(request *dto.LikeDeleteRequest) error {
	if err := s.validator.Struct(request); err != nil {
		s.logger.Warn().Err(err).Msg("failed to validate request")
		return err
	}
	totalPost, err := s.postRepository.CountById(request.PostId)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed count by id to database")
		return err
	}
	if totalPost < 1 {
		s.logger.Warn().Msg("post not found")
		return response.Except(404, "post not found")
	}
	totalUser, err := api.UserCountById(request.UserId)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed count by id to user service")
		return err
	}
	if totalUser < 1 {
		s.logger.Warn().Msg("user not found")
		return response.Except(404, "user not found")
	}
	like, err := s.likeRepository.FindByPostIdAndUserId(request.PostId, request.UserId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.Warn().Err(err).Msg("like not found")
			return response.Except(http.StatusNotFound, "like not found")
		}
		s.logger.Error().Err(err).Msg("failed find by post_id and user_id to database")
		return err
	}
	if err := s.postCache.DeleteById(request.PostId); err != nil {
		s.logger.Error().Err(err).Msg("failed delete by id to cache")
	}
	if err := s.likeRepository.Delete(like.Id); err != nil {
		s.logger.Error().Err(err).Msg("failed delete to database")
		return err
	}
	s.logger.Info().Str("like_id", strconv.Itoa(int(like.Id))).Msg("post unlike success")
	return nil
}

// comment
func (s *PostService) PostComment(request *dto.CommentAddRequest) error {
	if err := s.validator.Struct(request); err != nil {
		s.logger.Warn().Err(err).Msg("failed to validate request")
		return err
	}
	totalPost, err := s.postRepository.CountById(request.PostId)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed count by id to database")
		return err
	}
	if totalPost < 1 {
		s.logger.Warn().Msg("post not found")
		return response.Except(404, "post not found")
	}
	totalUser, err := api.UserCountById(request.UserId)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed count by id to user service")
		return err
	}
	if totalUser < 1 {
		s.logger.Warn().Msg("user not found")
		return response.Except(404, "user not found")
	}
	comment := &entity.Comment{
		PostId:      request.PostId,
		UserId:      request.UserId,
		Description: request.Description,
	}
	if err := s.postCache.DeleteById(request.PostId); err != nil {
		s.logger.Error().Err(err).Msg("failed delete by id to cache")
	}
	if err := s.commentRepository.Save(comment); err != nil {
		s.logger.Error().Err(err).Msg("failed save to database")
		return err
	}
	s.logger.Info().Str("post_id", strconv.Itoa(int(request.PostId))).Msg("post comment success")
	return nil
}
func (s *PostService) PostDeleteComment(request *dto.CommentDeleteRequest) error {
	if err := s.validator.Struct(request); err != nil {
		s.logger.Warn().Err(err).Msg("failed to validate request")
		return err
	}
	totalPost, err := s.postRepository.CountById(request.PostId)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed count by id to database")
		return err
	}
	if totalPost < 1 {
		s.logger.Warn().Msg("post not found")
		return response.Except(404, "post not found")
	}
	totalUser, err := api.UserCountById(request.UserId)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed count by id to user service")
		return err
	}
	if totalUser < 1 {
		s.logger.Warn().Msg("user not found")
		return response.Except(404, "user not found")
	}
	comment, err := s.commentRepository.FindByPostIdAndUserId(request.PostId, request.UserId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.Warn().Err(err).Msg("comment not found")
			return response.Except(http.StatusNotFound, "comment not found")
		}
		s.logger.Error().Err(err).Msg("failed find by post_id and user_id to database")
		return err
	}
	if err := s.postCache.DeleteById(request.PostId); err != nil {
		s.logger.Error().Err(err).Msg("failed delete by id to cache")
	}
	if err := s.commentRepository.Delete(comment.Id); err != nil {
		s.logger.Error().Err(err).Msg("failed delete to database")
		return err
	}
	s.logger.Info().Str("comment_id", strconv.Itoa(int(comment.Id))).Msg("post delete comment success")
	return nil
}
