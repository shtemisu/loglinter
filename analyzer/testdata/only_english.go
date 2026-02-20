package testdata

import (
	"errors"
	"log"
	"log/slog"

	"go.uber.org/zap"
)

func testEnglish() {
	log.Fatal("failed to connect server")
	zap.Error(errors.New("out of range"))
	var slice []int
	slog.Any("error", slice)
}
