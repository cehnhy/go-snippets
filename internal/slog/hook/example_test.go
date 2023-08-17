package hook_test

import (
	"context"
	"log/slog"
	"os"
	"time"

	"bou.ke/monkey"

	"github.com/cehnhy/go-snippets/internal/slog/hook"
)

func patchTime() {
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2023, 8, 16, 0, 0, 0, 0, time.UTC)
	})
}

func ExampleNewJSONHandler() {
	patchTime()

	hander := hook.NewJSONHandler(os.Stdout, nil, func(ctx context.Context, r *slog.Record) {
		if contextId, ok := ctx.Value("contextId").(string); ok {
			r.Add(slog.String("contextId", contextId))
		}
	})
	logger := slog.New(hander)

	ctx := context.WithValue(context.Background(), "contextId", "123")
	logger.InfoContext(ctx, "hello world")

	// Output:
	// {"time":"2023-08-16T00:00:00Z","level":"INFO","msg":"hello world","contextId":"123"}
}

func ExampleNewTextHandler() {
	patchTime()

	hander := hook.NewTextHandler(os.Stdout, nil, func(ctx context.Context, r *slog.Record) {
		if contextId, ok := ctx.Value("contextId").(string); ok {
			r.Add(slog.String("contextId", contextId))
		}
	})
	logger := slog.New(hander)

	ctx := context.WithValue(context.Background(), "contextId", "123")
	logger.InfoContext(ctx, "hello world")

	// Output:
	// time=2023-08-16T00:00:00.000Z level=INFO msg="hello world" contextId=123
}
