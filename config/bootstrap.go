package config

import (
	"post-management/delivery/rest/handler"
	"post-management/delivery/rest/middleware"
	"post-management/delivery/rest/routes"
	"post-management/internal/repository"
	"post-management/internal/service"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Bootstrap struct {
	DB        *gorm.DB
	Cache     *memcache.Client
	Logger    zerolog.Logger
	Validator *validator.Validate
	Router    *chi.Mux
	Ch        *amqp.Channel
}

func Initialize(deps *Bootstrap) {
	// publisher

	// cache

	// repository
	postRepo := repository.NewPostRepository(deps.DB)
	likeRepo := repository.NewLikeRepository(deps.DB)

	// service
	postServ := service.NewPostService(deps.Logger, deps.Validator, postRepo, likeRepo)

	// handler
	postHand := handler.NewPostHandler(postServ)

	// middleware
	deps.Router.Use(middleware.ErrorHandler)

	// routes
	r := routes.Router{
		Router:      deps.Router,
		PostHandler: postHand,
	}
	r.New()

	// subcriber
}
