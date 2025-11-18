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
func (h *PostHandler) PostUpdate(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	request := new(dto.PostUpdateRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		panic(response.Except(400, "failed to parse json"))
	}
	err := h.postService.PostUpdate(id, request)
	if err != nil {
		panic(err)
	}
	response.Success(w, 200, "OK", r.URL.Path)
}
func (h *PostHandler) PostGetById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	result, err := h.postService.PostGetById(id)
	if err != nil {
		panic(err)
	}
	response.Success(w, 200, result, r.URL.Path)
}
func (h *PostHandler) PostDelete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	err := h.postService.PostDelete(id)
	if err != nil {
		panic(err)
	}
	response.Success(w, 200, "OK", r.URL.Path)
}
func (h *PostHandler) PostLike(w http.ResponseWriter, r *http.Request) {
	request := new(dto.LikeAddRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		panic(response.Except(400, "failed to parse json"))
	}
	err := h.postService.PostLike(request)
	if err != nil {
		panic(err)
	}
	response.Success(w, 201, "OK", r.URL.Path)
}
func (h *PostHandler) PostUnlike(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	err := h.postService.PostUnlike(id)
	if err != nil {
		panic(err)
	}
	response.Success(w, 200, "OK", r.URL.Path)
}
