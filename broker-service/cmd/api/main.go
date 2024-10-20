package main

import (
	"log"
	"fmt"
	"net/http"
)

const webPort = "80"

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Startin broker service on port %s\n",  webPort)

	//define http server
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
		Handler: app.route(),
	}

	// start the server
	err:=srv.ListenAndServe()
	if err != nil{
		log.Panic(err)
	}
}