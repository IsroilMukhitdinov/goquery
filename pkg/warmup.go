package pkg

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ReadAnyData(url string, query string) (*goquery.Selection, error) {
	var result *goquery.Selection
	response, err := http.Get(url)
	if err != nil {
		return result, err
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return result, err
	}

	result = doc.Find(query)

	return result, nil

}
