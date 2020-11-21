package controllers

import (
	"net/http"
)

type pingController struct{}
type pingControllerInterface interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

var (
	PingController pingControllerInterface = &pingController{}
)

func (c *pingController) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong pong"))
}
