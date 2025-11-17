package dto

type ApiWebResponse[T any] struct {
	Data  *T      `json:"data,omitempty"`
	Error *string `json:"error,omitempty"`
	Path  string  `json:"path"`
}
type ApiUserCountResponse struct {
	Total int64 `json:"total"`
}
