package main

import (
	"log"
	"net/http"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	log.Println("starting server at http://localhost:3000/")
	log.Fatal(http.ListenAndServe(":3000", &Redirector{}))
}
