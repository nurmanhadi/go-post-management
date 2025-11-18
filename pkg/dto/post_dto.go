package dto

import "time"

type PostAddRequest struct {
	UserId      int64  `json:"user_id" validate:"required"`
	Description string `json:"description" validate:"required"`
}
type PostUpdateRequest struct {
	Description string `json:"description" validate:"required"`
}
type PostResponse struct {
	Id           int64             `json:"id"`
	Description  string            `json:"description"`
	TotalLike    int               `json:"total_like"`
	TotalComment int               `json:"total_comment"`
	User         UserResponse      `json:"user"`
	Comments     []CommentResponse `json:"comments"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
}
