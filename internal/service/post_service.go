package service

import (
	"net/http"
	"post-management/internal/entity"
	"post-management/internal/repository"
	"post-management/pkg/api"
	"post-management/pkg/dto"
	"post-management/pkg/response"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type PostService struct {
	logger         zerolog.Logger
	validator      *validator.Validate
	postRepository *repository.PostRepository
	likeRepository *repository.LikeRepository
}

func NewPostService(logger zerolog.Logger, validator *validator.Validate, postRepository *repository.PostRepository, likeRepository *repository.LikeRepository) *PostService {
	return &PostService{
		logger:         logger,
		validator:      validator,
		postRepository: postRepository,
		likeRepository: likeRepository,
	}
}
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
	post := &entity.Post{
		UserId:      request.UserId,
		Description: request.Description,
	}
	if err := s.postRepository.Save(post); err != nil {
		s.logger.Error().Err(err).Msg("failed save to database")
		return err
	}
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
	resp := &dto.PostResponse{
		Id:          post.Id,
		Description: post.Description,
		User: dto.UserResponse{
			Id:       user.Id,
			Username: user.Username,
			Name: dto.UserNameInfo{
				FirstName: user.Name.FirstName,
				Lastname:  user.Name.Lastname,
			},
			AvatarUrl: user.AvatarUrl,
		},
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
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
	if err := s.postRepository.Delete(post.Id); err != nil {
		s.logger.Error().Err(err).Msg("failed delete post to database")
		return err
	}
	s.logger.Info().Str("post_id", id).Msg("post delete success")
	return nil
}
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
	if err := s.likeRepository.Delete(like.Id); err != nil {
		s.logger.Error().Err(err).Msg("failed delete to database")
		return err
	}
	s.logger.Info().Str("like_id", strconv.Itoa(int(like.Id))).Msg("post unlike success")
	return nil
}
