package pkg

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ReadHTMLString() ([]string, error) {

	data := `
	<html lang="en">
	<body>
	<p>List of words</p>
	<ul>
		<li>dark</li>
		<li>smart</li>
		<li>war</li>
		<li>cloud</li>
		<li>park</li>
		<li>cup</li>
		<li>worm</li>
		<li>water</li>
		<li>rock</li>
		<li>warm</li>
		<li>peace</li>
	</ul>
	<footer>footer for words</footer>
	</body>
	</html>`

	var result []string

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data))
	if err != nil {
		return result, err
	}

	result = doc.Find("li").Map(func(i int, s *goquery.Selection) string {
		return fmt.Sprintf("[%d] %s", i+1, s.Text())
	})

	return result, nil
}
