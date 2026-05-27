package main

import "github.com/behnamdehghannejad/vendor/internal/app"

func main() {
	err := app.Start()
	if err != nil {
		return
	}
}
