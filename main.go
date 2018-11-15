package main

import (
	"fmt"
	"gopkg.in/macaron.v1"
	"log"
	"net/http"
)

func main() {
	m := macaron.Classic()
	m.SetDefaultCookieSecret("12336")
	m.Get("/set", myHandler)
	m.Get("/get",myhandler2)
	log.Println("server is runing...")
	log.Println(http.ListenAndServe("0.0.0.0:4000", m))
	m.Run()
}

func myHandler(ctx *macaron.Context) string  {
	ctx.SetSecureCookie("user","lida")
	return "set cookie success"
}

func myhandler2(ctx *macaron.Context) string {

	name, err := ctx.GetSecureCookie("user")
	if err != true {
		log.Println("get cookie fail")
	}

	 return name
}

func myhandler3(ctx *macaron.Context) string {
	return fmt.Sprintf("Num: %d", ctx.Data["Num"])
}
