package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gabrieleromanato/quiet_hn/hn"
)

func main() {
	// parse flags
	var port, numStories int

	flag.IntVar(&port, "port", 3000, "the port to start the web server on")
	flag.IntVar(&numStories, "num_stories", 30, "the number of top stories to display")
	flag.Parse()

	inMemoryCache := cache{
		items: []item{},
		ids:   []int{},
	}

	tpl := template.Must(template.ParseFiles("./index.gohtml"))

	http.HandleFunc("/", handler(numStories, tpl, inMemoryCache))

	// Start the server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func handler(numStories int, tpl *template.Template, c cache) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		start := time.Now()
		var client hn.Client

		cachedIDs := c.getIDs()
		cachedItems := c.getItems()

		var topIDs []int
		var stories []item

		if len(cachedIDs) == 0 {
			ids, err := client.TopItems()
			if err != nil {
				http.Error(w, "Failed to load top stories", http.StatusInternalServerError)
				return
			}
			c.addIDs(ids)
			topIDs = ids
		} else {
			topIDs = cachedIDs
		}

		if len(cachedItems) == 0 {
			for i, id := range topIDs {
				n := i + 1
				hnItem, err := getItem(id)
				if err != nil {
					continue
				}
				stories = append(stories, hnItem)
				if n >= numStories {
					break
				}
			}
			c.addItems(stories)
		} else {
			stories = cachedItems
		}

		data := templateData{
			Stories: stories,
			Time:    time.Now().Sub(start),
		}
		err := tpl.Execute(w, data)
		if err != nil {
			log.Fatal(err)
		}
	})
}

func isStoryLink(item item) bool {
	return item.Type == "story" && item.URL != ""
}

func parseHNItem(hnItem hn.Item) item {
	ret := item{Item: hnItem}
	url, err := url.Parse(ret.URL)
	if err == nil {
		ret.Host = strings.TrimPrefix(url.Hostname(), "www.")
	}
	return ret
}

func getItem(id int) (item, error) {
	var client hn.Client
	hnItem, err := client.GetItem(id)
	if err != nil {
		return item{}, err
	}
	itemParsed := parseHNItem(hnItem)
	if isStoryLink(itemParsed) {
		return itemParsed, nil
	}
	return item{}, fmt.Errorf("item %d is not a story link", id)
}

// item is the same as the hn.Item, but adds the Host field
type item struct {
	hn.Item
	Host string
}

type templateData struct {
	Stories []item
	Time    time.Duration
}

type cache struct {
	items []item
	ids   []int
}

func (c *cache) addIDs(ids []int) {
	c.ids = append(c.ids, ids...)
}

func (c *cache) getIDs() []int {
	return c.ids
}

func (c *cache) addItems(item []item) {
	c.items = append(c.items, item...)
}

func (c *cache) getItems() []item {
	return c.items
}
