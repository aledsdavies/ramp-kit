package routing

import (
	"bytes"
	"net/http"

	"github.com/a-h/templ"
	"github.com/asdavies/auth/internal/models"
	"github.com/asdavies/auth/views/layouts"
)

type pageBuilder func() templ.Component

func htmlRenderer(pageBuilder pageBuilder) http.HandlerFunc {
       return func(w http.ResponseWriter, r *http.Request) {
        var buf bytes.Buffer

        // Render the templ.Component to the buffer
        if err := pageBuilder().Render(r.Context(), &buf); err != nil {
            http.Error(w, "Error rendering page", http.StatusInternalServerError)
            return
        }

        // Get the rendered HTML as a string
        renderedPage := buf.String()

        // Render the layout with the rendered page HTML
        templ.Handler(layouts.Html(models.MetaInfo{}, renderedPage)).ServeHTTP(w, r)
    }
}
