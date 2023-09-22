package main

import (
	"log"
	"os"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func main() {
	pages := []string{"ex1.html", "ex2.html", "ex3.html", "ex4.html"}
	for _, page := range pages {
		links := LinkParser(page)
		for _, link := range links {
			log.Printf("href: %s, text: %s\n", link.Href, link.Text)
		}
	}
}

func LinkParser(file string) []Link {
	htmlFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer htmlFile.Close()
	doc, err := html.Parse(htmlFile)
	if err != nil {
		log.Fatal(err)
	}

	var f func(*html.Node)
	var links []Link
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			var link Link
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					link.Href = attr.Val
					break
				}
			}
			link.Text = n.FirstChild.Data
			links = append(links, link)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return links
}
