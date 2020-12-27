package main

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

// created seperate get method to avoid unnecessary repetition of code
func get(url string, client *http.Client) []byte {
	# create request from URL trimming off any whitespace from string
	req, _ := http.NewRequest("GET", strings.TrimSpace(url), nil)

	// set request headers to generate more human like traffic (I am not a robot?!)
	req.Header.Set("Connection","Keep-Alive")
	req.Header.Set("Accept-Language","en-US")
	req.Header.Set("User-Agent","Mozilla/5.0")

	resp, _ := client.Do(req)
	bytes, _ := ioutil.ReadAll(resp.Body)

	return bytes
}

func main() {
	var s SitemapIndex
	var n News
	client := &http.Client{}

	bytes := get("https://www.washingtonpost.com/news-sitemaps/index.xml", client)
	// call xml.Unmarshal method in main method to allow for different struct types each time
	xml.Unmarshal(bytes, &s)

	for _, Location := range s.Locations {
		bytes = get(Location, client)
		xml.Unmarshal(bytes, &n)
	}
}
