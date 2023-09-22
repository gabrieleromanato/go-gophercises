package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Stories map[string]Chapter

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

func main() {
	stories := getStories()

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler(stories))
	r.HandleFunc("/story/{name}", StoryHandler(stories))

	log.Fatal(http.ListenAndServe(":8000", r))

}

func getStories() Stories {
	var stories Stories
	jsonFile, err := os.Open("gopher.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	jsonParser := json.NewDecoder(jsonFile)
	jsonParser.Decode(&stories)
	return stories
}

func (s Stories) getChapter(name string) Chapter {
	ch, ok := s[name]
	if !ok {
		return Chapter{}
	}
	return ch
}

func HomeHandler(stories Stories) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		intro := stories.getChapter("intro")
		if intro.Title == "" {
			http.NotFound(w, r)
			return
		}
		t, err := template.ParseFiles("template/home.html")
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(w, intro)
	}
}

func StoryHandler(stories Stories) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		chapter := stories.getChapter(vars["name"])
		if chapter.Title == "" {
			http.NotFound(w, r)
			return
		}
		t, _ := template.ParseFiles("template/story.html")
		t.Execute(w, chapter)
	}
}
