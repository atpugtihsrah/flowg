package api

import (
	"context"
	"log/slog"

	"net/http"

	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"

	"link-society.com/flowg/internal/models"
	apiUtils "link-society.com/flowg/internal/utils/api"

	"link-society.com/flowg/internal/storage/auth"
	"link-society.com/flowg/internal/storage/config"
)

type BackupConfigRequest struct{}

type BackupConfigResponse struct {
	usecase.OutputWithEmbeddedWriter
}

func BackupConfigUsecase(
	authStorage *auth.Storage,
	configStorage *config.Storage,
) usecase.Interactor {
	u := usecase.NewInteractor(
		apiUtils.RequireScopesApiDecorator(
			authStorage,
			[]models.Scope{
				models.SCOPE_READ_PIPELINES,
				models.SCOPE_READ_TRANSFORMERS,
				models.SCOPE_READ_ALERTS,
			},
			func(
				ctx context.Context,
				req BackupConfigRequest,
				resp *BackupConfigResponse,
			) error {
				resp.Writer.(http.ResponseWriter).Header().Set("Content-Type", "application/octet-stream")
				resp.Writer.(http.ResponseWriter).Header().Set("Content-Disposition", "attachment; filename=config.db")
				resp.Writer.(http.ResponseWriter).Header().Set("Cache-Control", "no-cache")

				err := configStorage.Backup(ctx, resp.Writer)
				resp.Writer.(http.Flusher).Flush()
				if err != nil {
					slog.ErrorContext(
						ctx,
						"Failed to backup configuration database",
						slog.String("channel", "api"),
						slog.String("error", err.Error()),
					)

					return status.Wrap(err, status.Internal)
				}

				return nil
			},
		),
	)

	u.SetName("backup_config")
	u.SetTitle("Backup Configuration")
	u.SetDescription("Download a full snapshot of the configuration database.")
	u.SetTags("backup")

	u.SetExpectedErrors(status.Unauthenticated, status.PermissionDenied, status.Internal)

	return u
}
