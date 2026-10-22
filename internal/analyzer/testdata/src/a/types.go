package a

import (
	"context"
	l "log"
	"log/slog"
)

func testAliases() {
	l.Print("Alias") // want `log-message must be start lowercase`
	logger := slog.Default()
	logger.Info("Method")                             // want `log-message must be start lowercase`
	slog.InfoContext(context.Background(), "Context") // want `log-message must be start lowercase`
	slog.String("Key", "Value")
}
