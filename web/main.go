package web

import (
	"embed"

	"net/http"

	"link-society.com/flowg/internal/pipelines"
	"link-society.com/flowg/internal/storage"

	"link-society.com/flowg/web/controllers"
)

//go:embed static/**/*.css
//go:embed static/**/*.woff2
//go:embed static/**/*.js
var staticfiles embed.FS

//go:generate templ generate

func NewHandler(
	db *storage.Storage,
	pipelinesManager *pipelines.Manager,
) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("GET /static/", http.FileServer(http.FS(staticfiles)))

	mux.Handle("/dashboard/", controllers.DashboardController(db, pipelinesManager))

	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/dashboard/", http.StatusPermanentRedirect)
	})

	return mux
}
