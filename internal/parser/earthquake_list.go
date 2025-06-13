package parser

import (
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

func ParseEarthquakeListDoc(doc *goquery.Document) ([]string, error) {
	var earthquakeIndex []string
	doc.Find("#eqhist table tr").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			return
		}
		infoTd := s.Find("td").First()
		if infoTd.Length() == 0 {
			return
		}
		aTag := infoTd.Find("a").First()
		if aTag.Length() == 0 {
			return
		}
		href, exists := aTag.Attr("href")
		if !exists {
			return
		}
		indexRe := regexp.MustCompile(`(\d+)`)
		match := indexRe.FindStringSubmatch(href)
		if len(match) > 1 {
			earthquakeIndex = append(earthquakeIndex, match[1])
		}
	})
	// fmt.Println(earthquakeIndex)
	return earthquakeIndex, nil
}
