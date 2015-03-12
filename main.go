package main

import (
	"log"
	"net/http"

	cl "github.com/minhajuddin/config"
)

var config struct {
	Host             string `yaml:"host"`
	Port             string `yaml:"port"`
	ENV              string `yaml:"env"`
	DB               string `yaml:"db"`
	MaxDBConnections int    `yaml:"max_db_connections"`
}

func main() {
	//set log format
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	log.Println("starting server at http://localhost:3000/")

	//load config
	cl.LoadFromFile("./config.yml", &config, log.Println)

	log.Println(config)

	connectDB()

	//start http server
	http.HandleFunc(config.Host+"/", indexHandler)
	http.Handle("/", &Redirector{})

	log.Fatal(http.ListenAndServe(config.Port, nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
}
