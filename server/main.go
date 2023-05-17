package main

import (
	"net/http"
	"log"

	"github.com/fatih/color"
	middlewares "github.com/shangsuru/ctf-website/handlers"
	"github.com/shangsuru/ctf-website/routes"
	
)

func main() {
	port := middlewares.DotEnvVariable("PORT")
	color.Cyan("🌏 Server running on localhost:" + port)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	router := routes.Routes()

	http.ListenAndServe(":" + port, middlewares.LogRequest(router))
}

