package main

import (
	"sola-test-task/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}
