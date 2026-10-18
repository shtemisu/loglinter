package a

import (
	"log"
)

func testRussian() {
	log.Fatal("ошибка к подключению к серверу") // want `log-message must be in English`
	log.Println("ошибка извлечения данных")     // want `log-message must be in English`
}
