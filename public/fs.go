package public

import (
	"context"
	"crypto/md5"
	"embed"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"strings"

	"github.com/asdavies/auth/internal/assert"
	"github.com/asdavies/auth/internal/middleware"
	"github.com/go-chi/chi/v5"
)

//go:embed css/*.css
var stylesFS embed.FS

//go:embed css/styles.map.json
var stylesMapFile embed.FS

type fileVersions map[string]string

func (f fileVersions) ResourcePath(path string) string {
	hash, ok := f[path]
	assert.PanicIf(!ok, fmt.Sprintf("expected the file path '%s' to exist in FileVersions. Current FileVersions: %+v", path, f))
	return fmt.Sprintf("/styles/%s/%s", hash, path)
}

func (f fileVersions) BuildHandlers(r chi.Router) {
	// Serve the embedded static files
	cssDir, err := fs.Sub(stylesFS, "css")
	assert.PanicIfError(err, fmt.Sprintf("the application should have loaded the styles correctly. Error: %v", err))

	fileServer := http.FileServer(http.FS(cssDir))

	for path, hash := range f {
		url := fmt.Sprintf("/styles/%s/%s", hash, path)
		r.Handle(url, http.StripPrefix(fmt.Sprintf("/styles/%s/", hash), fileServer))
	}
}

type cssClassMap map[string]string

func (ccm cssClassMap) GetClassName(path string) (string, string) {
	class, ok := ccm[path]
	if !ok {
		keys := make([]string, 0, len(ccm))
		for k := range ccm {
			keys = append(keys, k)
		}
		assert.PanicIf(!ok, fmt.Sprintf("CSS value for '%s' does not exist. Available keys: %v", path, keys))
	}

	// Find the last index of '/'
	lastIndex := strings.LastIndex(path, ".")
	filePath := fmt.Sprintf("%s.min.css", strings.ReplaceAll(path[:lastIndex], ".", "/"))

	return class, filePath
}

type CSSLoader struct {
	cssModules      cssClassMap
	cssFileVersions fileVersions
	alwaysLoads     []string
	loadFiles       map[string][]string
}

// add ensures that the same CSS path is not added multiple times for the same request ID
func (cl *CSSLoader) add(ctx context.Context, path string) {
	requestID := middleware.FromContext(ctx)

	versionedPath := cl.cssFileVersions.ResourcePath(path)

    // Ensure the requestID entry exists in the map
    if cl.loadFiles[requestID] == nil {
        cl.loadFiles[requestID] = append([]string{}, cl.alwaysLoads...)
    }

	// Check if the path is already present
	for _, v := range cl.loadFiles[requestID] {
		if v == versionedPath {
			return
		}
	}

	// Add the new path
	cl.loadFiles[requestID] = append(cl.loadFiles[requestID], versionedPath)
}

// loadedStyles retrieves and returns the loaded styles, ensuring it's called only once per context
func (cl *CSSLoader) loadedStyles(ctx context.Context) []string {
	requestID := middleware.FromContext(ctx)
	usedFiles := cl.loadFiles[requestID]

	// Delete the requestID entry from the map
	delete(cl.loadFiles, requestID)

	return usedFiles
}

func RegisterGlobalStyles(paths ...string) {
	for _, path := range paths {
		cssLoader.alwaysLoads = append(cssLoader.alwaysLoads, cssLoader.cssFileVersions.ResourcePath(path))
	}
}

func CSS(ctx context.Context, path string) string {
	class, cssPath := cssLoader.cssModules.GetClassName(path)
	cssLoader.add(ctx, cssPath)
	return class
}

func LoadPageStyles(ctx context.Context) []string {
	return cssLoader.loadedStyles(ctx)
}

func BuildCssHandlers(r chi.Router) {
	cssLoader.cssFileVersions.BuildHandlers(r)
}

var cssLoader = &CSSLoader{
	cssModules:      make(cssClassMap),
	cssFileVersions: make(fileVersions),
	alwaysLoads:     []string{},
	loadFiles:       make(map[string][]string),
}

func init() {
	// Parse the JSON file containing CSS mappings
	data, err := stylesMapFile.ReadFile("css/styles.map.json")
	if err != nil {
		log.Fatalf("Failed to read CSS module data: %v", err)
	}

	if err := json.Unmarshal(data, &cssLoader.cssModules); err != nil {
		log.Fatalf("Failed to parse CSS module data: %v", err)
	}

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

			cssLoader.cssFileVersions[strings.TrimPrefix(path, "css/")] = hex.EncodeToString(hash.Sum(nil))
		}
		return nil
	})
}
