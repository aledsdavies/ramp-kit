package public

import (
	"crypto/md5"
	"embed"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"strings"

	"github.com/asdavies/auth/internal/assert"
	"github.com/go-chi/chi/v5"
)

//go:embed css/*.css
var stylesFS embed.FS

type FileVersions map[string]string


func (f FileVersions) ResourcePath(path string) string {
	hash, ok := f[path]
	assert.PanicIf(!ok, fmt.Sprintf("expected the file path '%s' to exist in FileVersions. Current FileVersions: %+v", path, f))
	return fmt.Sprintf("/styles/%s/%s", hash, path)
}

func (f FileVersions) BuildHandlers(r chi.Router) {
	// Serve the embedded static files
	cssDir, err := fs.Sub(stylesFS, "css")
	assert.PanicIfError(err, fmt.Sprintf("the application should have loaded the styles correctly. Error: %v", err))

	fileServer := http.FileServer(http.FS(cssDir))

	for path, hash := range f {
		url := fmt.Sprintf("/styles/%s/%s", hash, path)
		r.Handle(url, http.StripPrefix(fmt.Sprintf("/styles/%s/", hash), fileServer))
	}
}

var CssFileVersions = make(FileVersions)

func init() {
	// Walk through all files in the embedded file system
	fs.WalkDir(stylesFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			hash := md5.New()
			file, err := stylesFS.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			// Copy file content to hash
			if _, err := io.Copy(hash, file); err != nil {
				log.Fatal(err)
			}

			CssFileVersions[strings.TrimPrefix(path, "css/")] = hex.EncodeToString(hash.Sum(nil))
		}
		return nil
	})
}
