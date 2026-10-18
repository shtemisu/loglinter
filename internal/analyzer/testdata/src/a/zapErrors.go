package a

import (
	"errors"

	"go.uber.org/zap"
)

func testZap() {
	zap.Error(errors.New("failed to connect server"))
}
