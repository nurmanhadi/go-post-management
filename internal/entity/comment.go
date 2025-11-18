package entity

import "time"

type Comment struct {
	Id          int64
	PostId      int64
	UserId      int64
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Post        *Post
}
