package a

import "log"

func testSensetiveData() {
	api_key := "kyrlik"
	log.Fatal("api:", api_key) // want `log-message must not contains sensetive data`
	password := "password"
	log.Fatal("user_password:", password) // want `log-message must not contains sensetive data` `log-message must not contains spec chars`
	port := 8080
	log.Printf("server started on port: %d\n", port) // want `log-message must not contains spec chars`
}
