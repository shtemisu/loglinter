package testdata

import (
	"log"
	"log/slog"
)

func testSensetiveData() {
	api_key := "kyrlik"
	log.Fatal("api", api_key)
	password := "password"
	log.Fatal("user_password", password)
	slog.Error("author")
}
