package routes

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/orlovssky/gread/api"
	"github.com/orlovssky/gread/internal/controllers"
)

// routes - Setups chi router, middlewares and defines all api endpoints
func Routes() chi.Router {
	r := chi.NewRouter()

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Inject chi middleware
	// A good base middleware stack
	// Injects a request ID into the context of each request
	r.Use(middleware.RequestID)
	// Sets a http.Request's RemoteAddr to either X-Real-IP or X-Forwarded-For
	r.Use(middleware.RealIP)
	// Logs the start and end of each request with the elapsed processing time
	r.Use(middleware.Logger)
	// Gracefully absorb panics and prints the stack trace
	r.Use(middleware.Recoverer)
	// Sets HTTP response headers as content type JSON
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	// setup v1 subrouter
	r.Route("/v1", func(r chi.Router) {

		// health
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			api.JSON(w, 200, map[string]interface{}{"health_status": "online", "string": "test", "int": 3, "float": 1.32, "bool": true})
		})

		// auth
		r.Route("/auth", func(r chi.Router) {
			r.Post("/", controllers.Login) // POST /users
		})

		// // user
		// r.Route("/user", func(r chi.Router) {
		// 	r.Post("/", s.handleUserCreate)
		// 	r.Get("/", s.handleUserGet)
		// 	r.Put("/user/{id}", s.handleUserUpdate)
		// 	r.Delete("/user/{id}", s.handleUserDelete)
		// 	r.Get("/list", s.handleUserList)
		// })

	})

	return r
}