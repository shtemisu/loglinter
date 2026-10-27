package a

import (
	"errors"

	"go.uber.org/zap"
)

func testZap() {
	logger := zap.NewNop()
	sugar := logger.Sugar()
	err := errors.New("Bad")

	// *zap.Logger: сообщение — первый аргумент
	logger.Info("Connected")               // want `log-message must be start lowercase char`
	logger.Error("Failed", zap.Error(err)) // want `log-message must be start lowercase char`
	logger.Info("connected", zap.String("Key", "Value"))

	// *zap.SugaredLogger: Infof — format-строка, Infow — сообщение + пары ключ/значение
	sugar.Infof("Got %d items", 3) // want `log-message must be start lowercase char`
	sugar.Infof("value %-5d", 42)
	sugar.Infow("Request done", "status", 200) // want `log-message must be start lowercase char`

	// Конструкторы полей ничего не логируют — это не сообщения.
	zap.Error(errors.New("Bad"))
	zap.String("Key", "Value")
}
