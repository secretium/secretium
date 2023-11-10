package helpers

import (
	"embed"
	"io/fs"
	"log/slog"
	"net/http"
)

// StaticFileServerHandler handles a custom handler for serving static files from the embed static folder.
func StaticFileServerHandler(efs embed.FS) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the embedded static folder.
		static, err := fs.Sub(efs, "static_files")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			slog.Error(
				"failed to serve static file",
				"method", r.Method, "status", http.StatusBadRequest, "path", r.URL.Path,
				"error", err.Error(),
			)
			return
		}

		// Serve the file using the file server.
		http.FileServer(http.FS(static)).ServeHTTP(w, r)
	})
}
