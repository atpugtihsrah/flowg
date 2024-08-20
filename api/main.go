package api

import (
	"net/http"

	"github.com/swaggest/openapi-go/openapi31"
	"github.com/swaggest/rest/web"
	"github.com/swaggest/swgui/v5emb"

	"link-society.com/flowg/internal"
	"link-society.com/flowg/internal/pipelines"
	"link-society.com/flowg/internal/storage"
)

func NewHandler(
	db *storage.Storage,
	pipelinesManager *pipelines.Manager,
) http.Handler {
	service := web.NewService(openapi31.NewReflector())

	service.OpenAPISchema().SetTitle("Flowg API")
	service.OpenAPISchema().SetVersion(internal.FLOWG_VERSION)

	service.Docs("/api/docs", v5emb.New)

	service.Post("/api/v1/pipelines/{pipeline}/logs", IngestLogUsecase(pipelinesManager))

	service.Get("/api/v1/streams", ListStreamsUsecase(db))
	service.Get("/api/v1/streams/{stream}", QueryStreamUsecase(db))
	service.Get("/api/v1/streams/{stream}/fields", ListStreamFieldsUsecase(db))
	service.Delete("/api/v1/streams/{stream}", PurgeStreamUsecase(db))

	return service
}
