package pkg

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func FilterMore() ([]string, error) {
	var result []string

	response, err := http.Get("https://gowebexamples.com/")
	if err != nil {
		return result, err
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return result, err
	}

	doc.Find("a").FilterFunction(func(_ int, sel *goquery.Selection) bool {
		link, _ := sel.Attr("href")
		return isSecureLink(link)
	}).Each(func(_ int, sel *goquery.Selection) {
		result = append(result, sel.Text())
	})
	return result, nil
}

func isSecureLink(link string) bool {
	return strings.HasPrefix(link, "https")
}
