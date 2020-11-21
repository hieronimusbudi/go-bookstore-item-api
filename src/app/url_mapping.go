package app

import (
	"net/http"

	"github.com/hieronimusbudi/go-bookstore-item-api/src/controllers"
)

func mapUrls() {
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)
}
