package entity

import "time"

type Post struct {
	Id          int64
	UserId      int64
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Likes       []Like
}
