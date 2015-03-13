package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParseHost(t *testing.T) {
	for raw, clean := range map[string]string{
		"localhost:3000":  "localhost",
		"foo.lvh.me:3000": "foo.lvh.me",
		"lvh.me:3000":     "lvh.me",
		"localhost":       "localhost",
		"foo.lvh.me":      "foo.lvh.me",
		"lvh.ME":          "lvh.me",
		"LVH.me":          "lvh.me",
	} {
		host := parseHost(raw)
		if host != clean {
			t.Error(raw, clean, host, "invalid")
		}
	}
}

func TestInvertWWW(t *testing.T) {
	for from, to := range map[string]string{
		"localhost":     "www.localhost",
		"foo.lvh.me":    "www.foo.lvh.me",
		"www.lvh.me":    "lvh.me",
		"www.localhost": "localhost",
	} {
		//(23 * 2 - 8 + 4) * 30  + 75
		iwww := invertWWW(from)
		if to != iwww {
			t.Error(to, iwww, from, "invalid")
		}
	}
}

func TestRedirectorNonWWWtoNaked(t *testing.T) {
	red := &Redirector{}
	rec := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/about-us/contact?name=city&hyd=test", nil)
	req.Host = "mujju.com"
	log.Println(err)
	red.ServeHTTP(rec, req)

	if rec.Code != 301 || rec.Header().Get("Location") != "http://www.mujju.com/about-us/contact?name=city&hyd=test" {
		t.Error("invalid return code", rec)
	}
}

func TestRedirectorWWWtoNonWWWW(t *testing.T) {
	red := &Redirector{}
	rec := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/about-us/contact?name=city&hyd=test", nil)
	req.Host = "www.mujju.com"
	log.Println(err)
	red.ServeHTTP(rec, req)

	if rec.Code != 301 || rec.Header().Get("Location") != "http://mujju.com/about-us/contact?name=city&hyd=test" {
		t.Error("invalid return code", rec)
	}
}
