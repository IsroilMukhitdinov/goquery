package pkg

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Filter() ([]string, error) {
	var result []string
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

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data))
	if err != nil {
		return result, err
	}

	doc.Find("li").FilterFunction(func(_ int, s *goquery.Selection) bool {
		return strings.HasPrefix(s.Text(), "w")
	}).Each(func(_ int, s *goquery.Selection) {
		result = append(result, s.Text())
	})

	return result, nil
}
