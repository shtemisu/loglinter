package testdata

import (
	"log"
	"log/slog"
)

func testSpecChars() {
	log.Print("server started!🚀")
	log.Fatal("connection failed!!!")
	log.Fatalln("warning: something went wrong...")
	slog.Error("server started!🚀")
	slog.Error("connection failed!!!")
}
