package route

import (
	"net/http"

	"github.com/erwinwahyura/go-boilerplate/app/handler"
	"github.com/erwinwahyura/go-boilerplate/app/middleware"
	"github.com/erwinwahyura/go-boilerplate/app/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

// NewRoutes init Router
func NewRoutes(
	config model.Config,
	healthHandler handler.HealthHandler,
	userHandler handler.UserHandler,
	// another route here
) http.Handler {
	// Middleware
	mid := middleware.InitMiddleware(config)

	// Router
	r := chi.NewRouter()
	setMiddlewareGlobal(mid, r)

	// No Auth Routes
	r.Group(func(r chi.Router) {
		// Set Middleware

		// Static
		r.Get("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("static")) })

		// Swagger
		r.Get("/swagger/*", httpSwagger.WrapHandler)

		// Health Check
		r.Get("/healthcheck", healthHandler.Check)

		r.Route("/api/v1/public", func(r chi.Router) {
			// auth session
			r.Route("/auth", func(r chi.Router) {
				// for auth route here
			})

		})

		r.Route("/api/v1/user", func(r chi.Router) {
			r.Get("/create_user", userHandler.CreateUser)
		})
	})

	// Auth Routes
	r.Group(func(r chi.Router) {
		// Set Middleware
		r.Use(mid.Authenticate)
		// r.Use(mid.SetRequestID)
		// r.Use(mid.MiddlewareLogger)
		// r.Use(mid.ContextMandatoryRequest)
		r.Route("/api/v1/", func(r chi.Router) {
			r.Route("/users", func(r chi.Router) {

			})
		})
	})

	return r
}

// setMiddlewareGlobal set middleware global
func setMiddlewareGlobal(mid *middleware.GoMiddleware, r *chi.Mux) {
	// Logger
	r.Use(mid.LogRequest)

	// Cors
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Recovery
	r.Use(mid.RecoverPanic)
}
