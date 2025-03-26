package log

import (
	"context"
	"log/slog"
	"os"
)

const TraceID = "trace_id"

type MyHandler struct {
	slog.Handler
}

func (h MyHandler) Handle(ctx context.Context, r slog.Record) error {
	if traceID, ok := ctx.Value(TraceID).(string); ok {
		r.Add("trace_id", slog.StringValue(traceID))
	}

	return h.Handler.Handle(ctx, r)
}

func init() {

	var handler slog.Handler
	handler = slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
	})

	handler = &MyHandler{handler}

	slog.SetDefault(slog.New(handler))
}
