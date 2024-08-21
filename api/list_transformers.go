package api

import (
	"context"
	"log/slog"

	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"

	"link-society.com/flowg/internal/pipelines"
)

type ListTransformersRequest struct{}
type ListTransformersResponse struct {
	Success      bool     `json:"success"`
	Transformers []string `json:"transformers"`
}

func ListTransformersUsecase(pipelinesManager *pipelines.Manager) usecase.Interactor {
	u := usecase.NewInteractor(
		func(
			ctx context.Context,
			req ListTransformersRequest,
			resp *ListTransformersResponse,
		) error {
			transformers, err := pipelinesManager.ListTransformers()
			if err != nil {
				slog.ErrorContext(
					ctx,
					"Failed to list transformers",
					"channel", "api",
					"error", err.Error(),
				)

				resp.Success = false
				return status.Wrap(err, status.Internal)
			}

			resp.Success = true
			resp.Transformers = transformers

			return nil
		},
	)

	u.SetName("list_transformers")
	u.SetTitle("List Transformers")
	u.SetDescription("List Transformers")
	u.SetTags("transformers")

	u.SetExpectedErrors(status.Internal)

	return u
}
