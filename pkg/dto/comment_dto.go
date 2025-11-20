package dto

import "time"

type CommentAddRequest struct {
	PostId      int64  `json:"post_id" validate:"required"`
	UserId      int64  `json:"user_id" validate:"required"`
	Description string `json:"description" validate:"required"`
}
type CommentResponse struct {
	Id          int64        `json:"id"`
	Description string       `json:"description"`
	User        UserResponse `json:"user"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}
