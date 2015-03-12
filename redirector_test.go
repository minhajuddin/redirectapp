package main

import "testing"

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

func TestRedirector(t *testing.T) {
	//r := &Redirector{}

}
