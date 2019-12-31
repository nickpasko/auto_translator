package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"log"
	"io/ioutil"
	"github.com/PuerkitoBio/goquery"
	"bytes"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	search, err := template.ParseFiles("templates/search.html")
	if err != nil {
		fmt.Fprint(w, "Frag! Template not found!")
		log.Println("Frag! Template not found!")
		return
	}

	log.Println(r.URL)
	keys, ok := r.URL.Query()["search"]

	if !ok || len(keys[0]) < 1 {
		fmt.Fprint(w, "Url param 'search' is misssing")
		log.Println("Url param 'search' is misssing")
		return
	}
	//log.Println("search: " + keys[0])

	searchResult, err := getSearchResult(keys[0])
	if err != nil {
		log.Println("Failed to acquire search results")
		return
	}

	search.Execute(w, searchResult)
}

func getSearchResult(search string) (searchResult string, err error) {
	url := "https://catalog.tstarter.ru/detail-number/"+search
	log.Println("url: " + url)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Get request to tstarter.ru failed.")
		return "", err
	}
	log.Println("Response received")

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Response body read failed")
		return "", err
	}
	log.Println("Response body read success")
	log.Println(fmt.Sprintf("%s\n", body))

	dom, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		log.Println("Error creating DOM")
		return "", err
	}

	dom.Find(".search-by-id-results").Each(func(i int, s *goquery.Selection) {
		html , _ := goquery.OuterHtml(s)
		//log.Printf("%s\n", html)
		searchResult += fmt.Sprintf("%s", html)
	})

	return searchResult, nil
}
