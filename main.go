package main

import (
	"io/ioutil"
	"log"
	"net/http"

	cl "github.com/minhajuddin/config"
)

var (
	config struct {
		Host             string `yaml:"host"`
		Port             string `yaml:"port"`
		ENV              string `yaml:"env"`
		DB               string `yaml:"db"`
		MaxDBConnections int    `yaml:"max_db_connections"`
	}

	INDEX_HTML []byte
	err        error
)

func main() {
	//set log format
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	log.Println("starting server at http://localhost:3000/")

	//load config
	cl.LoadFromFile("./config.yml", &config, log.Println)

	log.Printf("loaded config: %+v\n", config)

	connectDB()

	INDEX_HTML, err = ioutil.ReadFile("./public/index.html")
	if err != nil {
		log.Println(err)
	}

	//start http server
	//index page is served by nginx
	log.Println(config.Host)
	http.HandleFunc(config.Host+"/", redirectsHandler)
	http.Handle("/", &Redirector{})

	log.Fatal(http.ListenAndServe(config.Port, nil))
}

func redirectsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write(INDEX_HTML)
	case "POST":
		createRedirectHandler(w, r)
	default:
		http.Error(w, "Invalid HTTP Verb. Only GET and POST are supported", 405)
	}
}

func createRedirectHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	createRedirect(r.PostForm)
}
