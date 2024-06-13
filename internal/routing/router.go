package routing

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"


	"github.com/asdavies/auth/views"
)

func NewRouter() *chi.Mux {
    r := chi.NewRouter()

    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        http.Redirect(w, r,  "/login", http.StatusSeeOther)
    })

    r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
        templ.Handler(views.Index("Aled")).ServeHTTP(w,r)
    })

    return r
}

