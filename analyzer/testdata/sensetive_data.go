package testdata

import "log"

func testSensetiveData() {
	api_key := "kyrlik"
	log.Fatal("api", api_key)
	password := "password"
	log.Fatal("user_password", password)
}
