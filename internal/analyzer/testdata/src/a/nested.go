package a

import "log"

// Вызовы логгера внутри других вызовов: анализатор должен найти их,
// даже если внешний вызов не принадлежит логгеру.

func runCallback(f func()) { f() }

func testNested() {
	defer func() {
		log.Print("In defer") // want `log-message must be start lowercase char`
	}()

	runCallback(func() {
		log.Print("In callback") // want `log-message must be start lowercase char`
	})

	go func() {
		log.Print("In goroutine") // want `log-message must be start lowercase char`
	}()
}
