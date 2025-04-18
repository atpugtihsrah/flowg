package logging

import (
	"log/slog"
	"net/http"
)

type middleware struct {
	handler http.Handler
}

var _ http.Handler = (*middleware)(nil)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewMiddleware(handler http.Handler) http.Handler {
	return &middleware{handler: handler}
}

func (m *middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	correlationId := r.Header.Get("X-Correlation-Id")
	ctx := WithCorrelationId(r.Context(), correlationId)
	req := r.WithContext(ctx)
	resp := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
	m.handler.ServeHTTP(resp, req)

	slog.InfoContext(
		req.Context(),
		"http request",
		slog.String("channel", "accesslog"),
		slog.Group("http",
			slog.Group("req",
				slog.String("method", req.Method),
				slog.String("url", req.URL.Path),
			),
			slog.Group("resp",
				slog.Int("status", resp.statusCode),
			),
		),
	)
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *responseWriter) Flush() {
	if f, ok := w.ResponseWriter.(http.Flusher); ok {
		f.Flush()
	}
}
