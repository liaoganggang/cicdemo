package main

import (
	"net/http"
	"testcicd/controllers"
	"testcicd/logger"
)

func main() {
	addr := ":9999"

	http.Handle("/webhook", &controllers.HookController{})
	if err := http.ListenAndServe(addr, nil); err != nil {
		logger.Error(err)
	}
}
