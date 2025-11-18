package dto

import "time"

type LikeAddRequest struct {
	PostId int64 `json:"post_id" validate:"required"`
	UserId int64 `json:"user_id" validate:"required"`
}
type LikeDeleteRequest struct {
	PostId int64 `json:"post_id" validate:"required"`
	UserId int64 `json:"user_id" validate:"required"`
}
type LikeResponse struct {
	Id        int64     `json:""`
	PostId    int64     `json:""`
	UserId    int64     `json:""`
	CreatedAt time.Time `json:""`
	UpdatedAt time.Time `json:""`
}
