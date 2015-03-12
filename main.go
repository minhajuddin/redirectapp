package main

import (
	"log"
	"net/http"
)

func main() {
	//set log format
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	log.Println("starting server at http://localhost:3000/")

	connectDB()

	//start http server
	log.Fatal(http.ListenAndServe(":3000", &Redirector{}))
}
