package dto

type LikeAddRequest struct {
	PostId int64 `json:"post_id" validate:"required"`
	UserId int64 `json:"user_id" validate:"required"`
}
