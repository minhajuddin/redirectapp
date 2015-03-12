package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"regexp"
	"strings"
)

var IGNORED_HOSTS = []string{"localhost", "redirectapp.com"}

type Redirector struct{}

func lookup(host string) string {
	dest := ""
	err := db.Get(&dest, "SELECT rules FROM redirects WHERE host = $1", host)
	if !noRows(err) {
		log.Println(err)
	}
	return dest
}

var wwwrx = regexp.MustCompile(`\Awww\.`)

func invertWWW(host string) string {
	if wwwrx.MatchString(host) {
		return wwwrx.ReplaceAllString(host, "")
	}
	return "www." + host
}

func (red *Redirector) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	code := http.StatusMovedPermanently

	//clean and parse the host
	host := parseHost(r.Host)

	//TODO if host is an IGNORED_HOST return

	//lookup to see if we have a record for this host
	dest := lookup(host)

	//if we haven't found a lookup redirect the user to an inverse www domain with the
	//path intact e.g. foobar.com/hello/ to www.foobar.com/hello/
	if dest == "" {
		dest = invertWWW(host)
	} else {
		//if we found one transform the incoming path to the destination host
	}

	//add scheme and path and query string
	dest = "http://" + dest + r.URL.Path + "?" + r.URL.RawQuery

	log.Printf("redirecting %s%s to %s\n", host, r.RequestURI, dest)

	//write the headers for the response
	w.Header().Set("Location", dest)
	w.WriteHeader(code)

	// RFC2616 recommends that a short note "SHOULD" be included in the
	// response because older user agents may not understand 301/307.
	// Shouldn't send the response for POST or HEAD; that leaves GET.
	if r.Method == "GET" {
		escapedDest := html.EscapeString(dest)
		note := "<a href=\"" + escapedDest + "\">Moved to " + escapedDest + "</a>.\n"
		fmt.Fprintln(w, note)
	}

}

func parseHost(host string) string {
	idx := strings.Index(host, ":")
	if idx < 0 {
		idx = len(host)
	}
	return strings.ToLower(host[0:idx])
}
