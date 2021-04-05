package main

import (
	"log"
	"net/http"
)

func main()  {
	const port = "8080"
	RegisterRoutes();

	done := make(chan bool)
	go http.ListenAndServe(":" + port, nil)
	log.Printf("Server started at port %v", port)
	<-done
}