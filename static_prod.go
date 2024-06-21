//go:build !dev
// +build !dev

package gothstarter

import (
	"embed"
	"net/http"
)

//go:embed public
var publicFS embed.FS

func Public() http.Handler {
	return http.FileServerFS(publicFS)
}
