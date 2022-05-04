package pkg

import (
	"bytes"
	"io/ioutil"

	// "regexp"

	"github.com/PuerkitoBio/goquery"
)

func ReadLocalFile(name string) (string, error) {
	var result string

	file, err := ioutil.ReadFile(name)
	if err != nil {
		return result, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(file))
	if err != nil {
		return result, err
	}
	result = doc.Find("h1, p").Text()

	// reg, err := regexp.Compile(`\s{2,}`)
	// if err != nil {
	// 	return result, err
	// }
	// result = reg.ReplaceAllString(result, "\n")

	return result, nil
}
