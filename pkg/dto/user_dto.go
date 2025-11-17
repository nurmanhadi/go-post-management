package dto

type UserResponse struct {
	Id        int64        `json:"id"`
	Username  string       `json:"username"`
	Name      UserNameInfo `json:"name"`
	AvatarUrl *string      `json:"avatar_url"`
}
type UserNameInfo struct {
	FirstName *string `json:"first_name"`
	Lastname  *string `json:"last_name"`
}
