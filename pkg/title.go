package pkg

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func GetTitle() (string, error) {

	var title string
	name := "http://webcode.me"
	response, err := http.Get(name)
	if err != nil {
		return title, err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		desc := fmt.Sprintf("non 200 status code recieved: %d %s", response.StatusCode, response.Status)
		err = errors.New(desc)
		return title, err
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return title, err
	}

	title = doc.Find("title").Text()

	return title, nil
}
