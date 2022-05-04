package pkg

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func UnionChildren() ([]string, error) {
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

	result = doc.Find("p:first-child, p:last-child").Union(doc.Find("p:nth-child(3), p:nth-child(2)")).Map(func(_ int, s *goquery.Selection) string {
		return s.Text()
	})

	return result, nil
}
