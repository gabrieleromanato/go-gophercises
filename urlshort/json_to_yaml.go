package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const (
	destFile = "shorturls.yml"
)

type ShortURL struct {
	Permalink string `json:"permalink"`
	ShortLink string `json:"shortlink"`
	Slug      string `json:"slug"`
}

type ShortURLs []ShortURL

func ConvertJSONToYAML() {
	jsonFile, err := os.Open("shorturls.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	data, _ := io.ReadAll(jsonFile)
	var shorturls ShortURLs
	json.Unmarshal(data, &shorturls)
	yamlString := ""
	for _, shorturl := range shorturls {
		yamlString += fmt.Sprintf("- permalink: %s\n  shortlink: %s\n  slug: %s\n", shorturl.Permalink, shorturl.ShortLink, shorturl.Slug)
	}
	err = os.WriteFile(destFile, []byte(yamlString), 0644)
	if err != nil {
		fmt.Println(err)
	}
}
