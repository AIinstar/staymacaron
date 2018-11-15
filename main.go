package main

import (
	"gopkg.in/macaron.v1"
	"log"
	"net/http"
)

func main() {
	m := macaron.Classic()
	m.Get("/", myHandler)
	log.Println("server is runing...")
	log.Println(http.ListenAndServe("0.0.0.0:4000", m))
	m.Run()
}

func myHandler(ctx *macaron.Context) string {
	return  "the request path is:" + ctx.Req.RequestURI
}