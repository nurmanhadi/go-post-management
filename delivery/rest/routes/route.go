package routes

import (
	"post-management/delivery/rest/handler"

	"github.com/go-chi/chi/v5"
)

type Router struct {
	Router      *chi.Mux
	PostHandler *handler.PostHandler
}

func (r *Router) New() {
	r.Router.Route("/api", func(api chi.Router) {
		api.Route("/posts", func(posts chi.Router) {
			posts.Post("/", r.PostHandler.PostCreate)
			posts.Put("/{id}", r.PostHandler.PostUpdate)
			posts.Get("/{id}", r.PostHandler.PostGetById)
			posts.Delete("/{id}", r.PostHandler.PostDelete)

			posts.Route("/likes", func(likes chi.Router) {
				likes.Post("/", r.PostHandler.PostLike)
				likes.Delete("/", r.PostHandler.PostUnlike)
			})
		})
	})
}
