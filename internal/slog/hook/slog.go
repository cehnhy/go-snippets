package hook

import (
	"context"
	"io"
	"log/slog"
)

// jsonHandler
type jsonHandler struct {
	*slog.JSONHandler
	hook func(context.Context, *slog.Record)
}

func NewJSONHandler(w io.Writer, opts *slog.HandlerOptions, hook func(context.Context, *slog.Record)) *jsonHandler {
	if hook == nil {
		hook = func(context.Context, *slog.Record) {}
	}

	return &jsonHandler{
		JSONHandler: slog.NewJSONHandler(w, opts),
		hook:        hook,
	}
}

func (h *jsonHandler) Handle(ctx context.Context, r slog.Record) error {
	h.hook(ctx, &r)
	return h.JSONHandler.Handle(ctx, r)
}

// textHandler
type textHandler struct {
	*slog.TextHandler
	hook func(context.Context, *slog.Record)
}

func NewTextHandler(w io.Writer, opts *slog.HandlerOptions, hook func(context.Context, *slog.Record)) *textHandler {
	if hook == nil {
		hook = func(context.Context, *slog.Record) {}
	}

	return &textHandler{
		TextHandler: slog.NewTextHandler(w, opts),
		hook:        hook,
	}
}

func (h *textHandler) Handle(ctx context.Context, r slog.Record) error {
	h.hook(ctx, &r)
	return h.TextHandler.Handle(ctx, r)
}
