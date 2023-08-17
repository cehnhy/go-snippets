package setup_test

import (
	"log"
	"log/slog"
	"os"
	"time"

	"bou.ke/monkey"

	"github.com/cehnhy/go-snippets/internal/slog/setup"
)

func patchTime() {
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2023, 8, 16, 0, 0, 0, 0, time.UTC)
	})
}

func ExampleSetJSON() {
	patchTime()

	setup.SetJSON("go-snippets", "info", os.Stdout)
	log.Println("println message")
	slog.Debug("debug message") // no output
	slog.Info("info message")
	slog.Warn("warn message", "key", "value")
	slog.Error("error message", slog.String("key", "value"))

	// Output:
	// {"time":"2023-08-16T00:00:00Z","level":"INFO","msg":"println message"}
	// {"time":"2023-08-16T00:00:00Z","level":"INFO","msg":"info message"}
	// {"time":"2023-08-16T00:00:00Z","level":"WARN","msg":"warn message","go-snippets":{"key":"value"}}
	// {"time":"2023-08-16T00:00:00Z","level":"ERROR","msg":"error message","go-snippets":{"key":"value"}}
}

func ExampleSetText() {
	patchTime()

	setup.SetText("go-snippets", "info", os.Stdout)
	log.Println("println message")
	slog.Debug("debug message") // no output
	slog.Info("info message")
	slog.Warn("warn message", "key", "value")
	slog.Error("error message", slog.String("key", "value"))

	// Output:
	// time=2023-08-16T00:00:00.000Z level=INFO msg="println message"
	// time=2023-08-16T00:00:00.000Z level=INFO msg="info message"
	// time=2023-08-16T00:00:00.000Z level=WARN msg="warn message" go-snippets.key=value
	// time=2023-08-16T00:00:00.000Z level=ERROR msg="error message" go-snippets.key=value
}

func ExampleWithRolling() {
	patchTime()

	w := setup.WithRolling("go-snippets.log", os.Stdout)
	setup.SetText("go-snippets", "info", w)
	log.Println("println message")
	slog.Debug("debug message") // no output
	slog.Info("info message")
	slog.Warn("warn message", "key", "value")
	slog.Error("error message", slog.String("key", "value"))

	// Output:
	// time=2023-08-16T00:00:00.000Z level=INFO msg="println message"
	// time=2023-08-16T00:00:00.000Z level=INFO msg="info message"
	// time=2023-08-16T00:00:00.000Z level=WARN msg="warn message" go-snippets.key=value
	// time=2023-08-16T00:00:00.000Z level=ERROR msg="error message" go-snippets.key=value
}
