package entity

import "time"

type Like struct {
	Id        int64
	PostId    int64
	UserId    int64
	CreatedAt time.Time
	UpdatedAt time.Time
	Post      *Post
}
