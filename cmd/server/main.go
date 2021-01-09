package main

import (
	httpAPI "github.com/amburskui/tasklauncher/internal/app/interfaces/http"
)

func main() {
	api := httpAPI.NewAPI()
	api.ListenAndServe(":8000")
}
