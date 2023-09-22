package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	website := flag.String("website", "https://www.google.com", "website to parse")
	flag.Parse()
	link := *website
	if !isValidURL(link) {
		log.Fatal("Invalid URL")
	}
	foundUrls := make(map[string]bool)
	chUrls := make(chan string)
	chFinished := make(chan bool)
	seedUrls := []string{link}

	for _, url := range seedUrls {
		go linkParser(url, link, chUrls, chFinished)
	}

	for c := 0; c < len(seedUrls); {
		select {
		case url := <-chUrls:
			foundUrls[url] = true
			if !inSlice(url, seedUrls) {
				seedUrls = append(seedUrls, url)
				go linkParser(url, link, chUrls, chFinished)
			}
		case <-chFinished:
			c++
		}
	}

	fmt.Println("Found", len(foundUrls), "unique urls:")

	for url, _ := range foundUrls {
		fmt.Println(" - " + url)
	}
	createSiteMap(foundUrls)

	close(chUrls)

}

func isValidURL(u string) bool {
	_, err := url.Parse(u)
	if err != nil {
		return false
	}
	return true
}

func getHref(t html.Token) (ok bool, href string) {
	for _, attr := range t.Attr {
		if attr.Key == "href" {
			href = attr.Val
			ok = true
		}
	}
	return
}

func inSlice(s string, sl []string) bool {
	for _, a := range sl {
		if a == s {
			return true
		}
	}
	return false
}

func linkParser(url string, base string, ch chan string, finished chan bool) {
	fmt.Println("Found", url)
	htmlFile, err := http.Get(url)
	defer func() {
		finished <- true
	}()
	if err != nil {
		log.Fatal(err)
	}
	defer htmlFile.Body.Close()

	tokenizer := html.NewTokenizer(htmlFile.Body)
	for {
		tt := tokenizer.Next()
		switch {
		case tt == html.ErrorToken:
			return
		case tt == html.StartTagToken:
			t := tokenizer.Token()
			isAnchor := t.Data == "a"
			if !isAnchor {
				continue
			}
			ok, url := getHref(t)
			if !ok {
				continue
			}
			hasBase := strings.Index(url, base) == 0
			if hasBase {
				ch <- url
			}
		}
	}

}

func createSiteMap(links map[string]bool) {

	type Url struct {
		Loc string `xml:"loc"`
	}
	type UrlSet struct {
		XMLName xml.Name `xml:"urlset"`
		Xmlns   string   `xml:"xmlns,attr"`
		Urls    []Url    `xml:"url"`
	}
	var urls []Url
	for link, _ := range links {
		urls = append(urls, Url{Loc: link})
	}

	urlSet := UrlSet{Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9", Urls: urls}
	output, err := xml.MarshalIndent(urlSet, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	err = os.WriteFile("sitemap.xml", []byte(xml.Header+string(output)), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
