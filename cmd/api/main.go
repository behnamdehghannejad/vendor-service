package main

import "order-service/internal/app"

func main() {
	err := app.Start()
	if err != nil {
		return
	}
}
