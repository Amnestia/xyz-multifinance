package router

import (
	"net/http"

	"github.com/amnestia/xyz-multifinance/internal/config"
	"github.com/amnestia/xyz-multifinance/internal/domain/controller"
	intmiddleware "github.com/amnestia/xyz-multifinance/internal/middleware"
	"github.com/amnestia/xyz-multifinance/internal/middleware/auth"
	"github.com/amnestia/xyz-multifinance/pkg/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// Router struct containing router wrapper
type Router struct {
	Handler    chi.Router
	Controller *controller.Controller
	Config     config.Config
	Options    Options
	Auth       auth.AuthorizationModule
}

type Options struct{}

func New(opt Options, cfg config.Config, auth auth.AuthorizationModule, controller *controller.Controller) *Router {
	r := &Router{
		Handler:    chi.NewRouter(),
		Options:    opt,
		Config:     cfg,
		Auth:       auth,
		Controller: controller,
	}
	r.setMiddleware()
	r.registerRoute()
	err := chi.Walk(r.Handler, func(method, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		logger.Logger.Info().Msgf("%s : %s", method, route)
		return nil
	})
	if err != nil {
		logger.Logger.Err(logger.ErrorWrap(err, "NewRouter.chiWalk")).Send()
	}

	return r
}

// setMiddleware set default middleware
func (r *Router) setMiddleware() {
	r.Handler.Use(cors.Handler(cors.Options{
		AllowedOrigins:   r.Config.Server.Origin,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"User-Agent", "Content-Type", "Accept", "Accept-Encoding", "Accept-Language", "Cache-Control", "Connection", "DNT", "Host", "Origin", "Pragma", "Referer"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	r.Handler.Use(intmiddleware.PanicRecovery)
	r.Handler.Use(middleware.CleanPath)
	r.Handler.Use(middleware.RealIP)
	r.Handler.Use(middleware.RequestID)
	r.Handler.Use(intmiddleware.Logger)
}

// registerRoute register API routes
func (router *Router) registerRoute() {
	router.Handler.Get("/ping", router.Controller.PingHandler.Ping)
	router.Handler.Post("/register", router.Controller.AuthHandler.Register)
	router.Handler.Post("/login", router.Controller.AuthHandler.Auth)

	router.Handler.Group(func(r chi.Router) {
		r.Use(router.Auth.Authorize)
		r.Get("/pingauth", router.Controller.PingHandler.Ping)
		r.Post("/partner/register", router.Controller.AuthHandler.RegisterNewPartner)
	})

	router.Handler.Group(func(r chi.Router) {
		r.Use(router.Auth.AuthorizeAPIKey)
		r.Post("/transaction", router.Controller.TransactionHandler.CreateNewTransaction)
	})

}
