package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/swaggest/openapi-go"
	"github.com/swaggest/openapi-go/openapi31"
	"github.com/swaggest/rest/nethttp"
	"github.com/swaggest/rest/web"
	"github.com/swaggest/swgui/v5emb"
	"github.com/swaggest/usecase"

	"link-society.com/flowg/internal/app"
	"link-society.com/flowg/internal/data/auth"
	"link-society.com/flowg/internal/data/lognotify"
	"link-society.com/flowg/internal/data/logstorage"
	"link-society.com/flowg/internal/data/pipelines"
)

func NewHandler(
	authDb *auth.Database,
	logDb *logstorage.Storage,
	pipelinesManager *pipelines.Manager,
	logNotifier *lognotify.LogNotifier,
) http.Handler {
	reflector := openapi31.NewReflector()
	service := web.NewService(reflector)

	service.OpenAPISchema().SetTitle("Flowg API")
	service.OpenAPISchema().SetVersion(app.FLOWG_VERSION)
	service.Docs("/api/docs", v5emb.New)

	service.With(
		nethttp.HTTPBearerSecurityMiddleware(
			service.OpenAPICollector,
			"tokenAuth",
			"Authentication using Personal Access Token",
			"PAT",
		),
		auth.ApiMiddleware(authDb),
	).Group(func(router chi.Router) {
		r := &routerWrapper{Router: router}

		r.Get("/api/v1/transformers", ListTransformersUsecase(authDb, pipelinesManager))
		r.Get("/api/v1/transformers/{transformer}", GetTransformerUsecase(authDb, pipelinesManager))
		r.Put("/api/v1/transformers/{transformer}", SaveTransformerUsecase(authDb, pipelinesManager))
		r.Delete("/api/v1/transformers/{transformer}", DeleteTransformerUsecase(authDb, pipelinesManager))
		r.Post("/api/v1/transformers/{transformer}/test", TestTransformerUsecase(authDb, pipelinesManager))

		r.Get("/api/v1/pipelines", ListPipelinesUsecase(authDb, pipelinesManager))
		r.Get("/api/v1/pipelines/{pipeline}", GetPipelineUsecase(authDb, pipelinesManager))
		r.Put("/api/v1/pipelines/{pipeline}", SavePipelineUsecase(authDb, pipelinesManager))
		r.Delete("/api/v1/pipelines/{pipeline}", DeletePipelineUsecase(authDb, pipelinesManager))
		r.Post("/api/v1/pipelines/{pipeline}/logs", IngestLogUsecase(authDb, pipelinesManager))

		r.Get("/api/v1/streams", ListStreamsUsecase(authDb, logDb))
		r.Get("/api/v1/streams/{stream}", GetStreamUsecase(authDb, logDb))
		r.Put("/api/v1/streams/{stream}", ConfigureStreamUsecase(authDb, logDb))
		r.Get("/api/v1/streams/{stream}/logs", QueryStreamUsecase(authDb, logDb))
		r.Get("/api/v1/streams/{stream}/fields", ListStreamFieldsUsecase(authDb, logDb))
		r.Delete("/api/v1/streams/{stream}", PurgeStreamUsecase(authDb, logDb))

		service.OpenAPICollector.AnnotateOperation(
			"GET", "/api/v1/streams/{stream}/logs/watch",
			func(oc openapi.OperationContext) error {
				contentUnits := oc.Response()
				for i, cu := range contentUnits {
					if cu.HTTPStatus == 200 {
						cu.ContentType = "text/event-stream"
						cu.Description = "Stream of log entries"
						cu.Format = "Server-Sent Events"
					}

					contentUnits[i] = cu
				}

				return nil
			},
		)
		r.Get("/api/v1/streams/{stream}/logs/watch", WatchLogsUsecase(authDb, logNotifier))

		r.Get("/api/v1/roles", ListRolesUsecase(authDb))
		r.Put("/api/v1/roles/{role}", SaveRoleUsecase(authDb))
		r.Delete("/api/v1/roles/{role}", DeleteRoleUsecase(authDb))

		r.Get("/api/v1/users", ListUsersUsecase(authDb))
		r.Put("/api/v1/users/{user}", SaveUserUsecase(authDb))
		r.Delete("/api/v1/users/{user}", DeleteUserUsecase(authDb))

		r.Get("/api/v1/tokens", ListTokensUsecase(authDb))
		r.Post("/api/v1/token", CreateTokenUsecase(authDb))
		r.Delete("/api/v1/tokens/{token-uuid}", DeleteTokenUsecase(authDb))
	})

	return service
}

type routerWrapper struct {
	chi.Router
}

func (r *routerWrapper) Get(pattern string, u usecase.Interactor) {
	r.Method(http.MethodGet, pattern, nethttp.NewHandler(u))
}

func (r *routerWrapper) Post(pattern string, u usecase.Interactor) {
	r.Method(http.MethodPost, pattern, nethttp.NewHandler(u))
}

func (r *routerWrapper) Put(pattern string, u usecase.Interactor) {
	r.Method(http.MethodPut, pattern, nethttp.NewHandler(u))
}

func (r *routerWrapper) Delete(pattern string, u usecase.Interactor) {
	r.Method(http.MethodDelete, pattern, nethttp.NewHandler(u))
}
