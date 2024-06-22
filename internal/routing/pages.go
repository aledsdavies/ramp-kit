package routing

import (
	"github.com/a-h/templ"
	"github.com/asdavies/auth/views"
	"github.com/go-chi/chi/v5"
)


func pages() chi.Router {
	r := chi.NewRouter()


	r.Group(func(r chi.Router) {
		r.Get("/login", htmlRenderer(func() templ.Component {
			return views.Index("Aled")
		}))
	})

    return r
}
