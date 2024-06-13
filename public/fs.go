package public

import (
	"crypto/md5"
	"embed"
	"encoding/hex"
	"io"
	"io/fs"
	"log"
)

//go:embed *.css
var StylesFS embed.FS
var CssVersion string

func init() {
	// Initialize an MD5 hash object
	hash := md5.New()

	// Walk through all files in the embedded file system
	fs.WalkDir(StylesFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			file, err := StylesFS.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			// Copy file content to hash
			if _, err := io.Copy(hash, file); err != nil {
				log.Fatal(err)
			}
		}
		return nil
	})

	// Compute final hash value and convert to string
	CssVersion = hex.EncodeToString(hash.Sum(nil))
}
