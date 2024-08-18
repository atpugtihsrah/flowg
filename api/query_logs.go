package api

import (
	"context"
	"log/slog"

	"time"

	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"

	"link-society.com/flowg/internal/filterdsl"
	"link-society.com/flowg/internal/storage"
)

type QueryStreamRequest struct {
	Stream string    `path:"stream" minLength:"1"`
	From   time.Time `query:"from" format:"date-time" required:"true"`
	To     time.Time `query:"to" format:"date-time" required:"true"`
	Filter *string   `query:"filter"`
}

type QueryStreamResponse struct {
	Success bool               `json:"success"`
	Records []storage.LogEntry `json:"records"`
}

func QueryStreamUsecase(db *storage.Storage) usecase.Interactor {
	u := usecase.NewInteractor(
		func(
			ctx context.Context,
			req QueryStreamRequest,
			resp *QueryStreamResponse,
		) error {
			var filter storage.Filter

			if req.Filter != nil {
				var err error
				filter, err = filterdsl.Compile(*req.Filter)
				if err != nil {
					slog.ErrorContext(
						ctx,
						"Failed to compile filter",
						"channel", "api",
						"stream", req.Stream,
						"error", err.Error(),
					)

					resp.Success = false
					resp.Records = nil
					return status.Wrap(err, status.InvalidArgument)
				}
			} else {
				filter = nil
			}

			records, err := db.Query(req.Stream, req.From, req.To, filter)
			if err != nil {
				slog.ErrorContext(
					ctx,
					"Failed to query logs",
					"stream", req.Stream,
					"error", err.Error(),
				)
				return status.Wrap(err, status.Internal)
			}

			resp.Success = true
			resp.Records = records
			return nil
		},
	)

	u.SetName("query_stream")
	u.SetTitle("Query Stream")
	u.SetDescription("Query logs from a stream")
	u.SetTags("streams")

	u.SetExpectedErrors(status.InvalidArgument, status.Internal)

	return u
}
