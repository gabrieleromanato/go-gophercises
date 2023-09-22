package main

import (
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

type ShortURI struct {
	Permalink string `yaml:"permalink"`
	ShortLink string `yaml:"shortlink"`
	Slug      string `yaml:"slug"`
}

type ShortURIs []ShortURI

func main() {
	shortlinks := parseYML()
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		shortl := r.URL.Path[1:]
		permalink := findShortLink(shortlinks, shortl)
		if permalink == "" {
			http.NotFound(w, r)
			return
		}
		http.Redirect(w, r, permalink, http.StatusMovedPermanently)
	})
	http.ListenAndServe(":8080", mux)
}

func parseYML() []ShortURI {
	file, err := os.ReadFile("shorturls.yml")
	if err != nil {
		panic(err)
	}
	var shortlinks ShortURIs
	err = yaml.Unmarshal(file, &shortlinks)
	if err != nil {
		panic(err)
	}
	return shortlinks
}

func findShortLink(shortlinks []ShortURI, shortl string) string {
	for _, shortlink := range shortlinks {
		if shortlink.ShortLink == shortl {
			return shortlink.Permalink
		}
	}
	return ""
}
