package jpquakeparser

import (
	"earthquake-crawler/internal/config"
	"earthquake-crawler/internal/util"
	"regexp"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func ParseJapanEarthquakeListDoc(doc *goquery.Document) ([]string, error) {
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
	var resEqtList []string
	for _, eqt := range earthquakeIndex {
		t, err := time.ParseInLocation("20060102150405", eqt, util.GetTokyoLocation())
		if err != nil {
			return nil, err
		}
		if util.GetCurrentJapanTime().Sub(t) >= time.Duration(config.Cfg.JPQuake.ParseAfterMinute)*time.Minute {
			resEqtList = append(resEqtList, eqt)
		}
	}
	return resEqtList, nil
}
