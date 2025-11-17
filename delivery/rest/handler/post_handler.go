package handler

import (
	"net/http"
	"post-management/internal/service"
	"post-management/pkg/dto"
	"post-management/pkg/response"

	"github.com/goccy/go-json"
)

type PostHandler struct {
	postService *service.PostService
}

func NewPostHandler(postService *service.PostService) *PostHandler {
	return &PostHandler{
		postService: postService,
	}
}
func (h *PostHandler) PostCreate(w http.ResponseWriter, r *http.Request) {
	request := new(dto.PostAddRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		panic(response.Except(400, "failed to parse json"))
	}
	err := h.postService.PostCreate(request)
	if err != nil {
		panic(err)
	}
	response.Success(w, 201, "OK", r.URL.Path)
}
