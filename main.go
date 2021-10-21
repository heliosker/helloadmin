package main

import (
	"fmt"
	"helloadmin/routers"
	"log"
	"net/http"
)

func main() {
	r := routers.InitRouter()
	//port, _ := config.Load().Section("server").Key("HTTP_PORT").Int()
	port := 9010
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        r,
		MaxHeaderBytes: 1 << 20,
	}

	e := s.ListenAndServe()
	if e != nil {
		log.Printf("Server error: %v", e)
	}
}
