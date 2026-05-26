package main

import "vendor-service/internal/app"

func main() {
	err := app.Start()
	if err != nil {
		return
	}
}
