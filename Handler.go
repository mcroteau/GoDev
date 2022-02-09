package main

import "net/http"

type Handler struct{}

func (handler *Handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	println(req.Method)
	resp.Write([]byte("This is my home page"))
}
