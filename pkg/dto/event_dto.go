package dto

import "time"

type EventProducer[T any] struct {
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
	Data      T         `json:"data"`
}
type EventPostCreatedProducer struct {
	PostId       int64     `json:"post_id"`
	UserId       int64     `json:"user_id"`
	Description  string    `json:"description"`
	TotalLike    int       `json:"total_like"`
	TotalComment int       `json:"total_comment"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
type EventPostUpdatedProducer struct {
	PostId      int64     `json:"post_id"`
	Description string    `json:"description"`
	UpdatedAt   time.Time `json:"updated_at"`
}
type EventLikeTotalProducer struct {
	PostId int64 `json:"post_id"`
	Total  bool  `json:"total"`
}
type EventCommentTotalProducer struct {
	PostId int64 `json:"post_id"`
	Total  bool  `json:"total"`
}
