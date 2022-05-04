package pkg

import (
	"errors"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func GetRakuQuestions() ([]string, error) {
	var result []string

	response, err := http.Get("https://stackoverflow.com/questions/tagged/raku")
	if err != nil {
		return result, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return result, errors.New("non 200 status code received")
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return result, err
	}

	doc.Find(".s-post-summary--content-title a").Each(func(_ int, s *goquery.Selection) {
		result = append(result, s.Text())
	})

	return result, nil
}
