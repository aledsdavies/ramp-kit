package routing

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	m "github.com/asdavies/auth/internal/middleware"
	"github.com/asdavies/auth/public"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(m.RequestIDMiddleware)

	r.Group(func(r chi.Router) {
		r.Use(middleware.Compress(9))
		r.Use(cacheControlMiddleware)

		public.BuildCssHandlers(r)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	})

    r.Mount("/", pages())

	return r
}

func cacheControlMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=31536000")
		next.ServeHTTP(w, r)
	})
}
