package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"helloadmin/routers"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	r := routers.InitRouter()
	port, _ := strconv.Atoi(os.Getenv("HTTP_PORT"))
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
