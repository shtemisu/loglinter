package a

import (
	"log"
	"log/slog"
)

func testSpecChars() {
	log.Print("server started!🚀")                   // want `log-message must not contains spec chars`
	log.Fatal("connection failed!!!")               // want `log-message must not contains spec chars`
	log.Fatalln("warning: something went wrong...") // want `log-message must not contains spec chars`
	slog.Error("server started!🚀")                  // want `log-message must not contains spec chars`
	slog.Error("connection failed!!!")              // want `log-message must not contains spec chars`
}
