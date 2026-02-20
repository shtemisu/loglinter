package testdata

import (
	"errors"
	"log"
	"log/slog"

	"go.uber.org/zap"
)

func russianLog() {
	log.Fatalf("ОШИБКА!")
	slog.Error("ошибка к подключению")
	zap.Error(errors.New("порт занят"))
}
