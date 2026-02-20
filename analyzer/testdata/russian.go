package testdata

import (
	"errors"

	"go.uber.org/zap"
)

func russianLog() {
	//log.Fatalf("ОШИБКА!")
	//slog.Error("ошибка к подключению")
	zap.Error(errors.New("порт занят"))
	err := errors.New("Ошибка к подключению к БД")
	zap.Error(err)
}
