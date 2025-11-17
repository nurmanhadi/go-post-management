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
	Id          int64        `json:"id"`
	Description string       `json:"description"`
	User        UserResponse `json:"user"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}
