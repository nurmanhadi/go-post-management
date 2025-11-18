package dto

type ApiWebPayload[T any] struct {
	Data  *T      `json:"data,omitempty"`
	Error *string `json:"error,omitempty"`
	Path  string  `json:"path"`
}
type ApiUserCountPayload struct {
	Total int64 `json:"total"`
}
type ApiUserPayload struct {
	Id        int64        `json:"id"`
	Username  string       `json:"username"`
	Name      UserNameInfo `json:"name"`
	AvatarUrl *string      `json:"avatar_url"`
}
type ApiUserNameInfoPayload struct {
	FirstName *string `json:"first_name"`
	Lastname  *string `json:"last_name"`
}
type ApiUserGetBySliceIdBody struct {
	Ids []int64 `json:"ids"`
}
