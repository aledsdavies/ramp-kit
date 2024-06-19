package routing

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/asdavies/auth/public"
	"github.com/asdavies/auth/views"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Group(func(r chi.Router) {
		r.Use(middleware.Compress(9))
        r.Use(cacheControlMiddleware)

        public.BuildCssHandlers(r)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	})

	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(views.Index("Aled")).ServeHTTP(w, r)
	})

	return r
}

func cacheControlMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Cache-Control", "public, max-age=31536000")
        next.ServeHTTP(w, r)
    })
}
