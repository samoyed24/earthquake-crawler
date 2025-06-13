package spider

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func GetEarthquakeListDoc() (*goquery.Document, error) {
	client := &http.Client{ // 超时时间
		Timeout: 5 * time.Second,
	}
	URL := "https://typhoon.yahoo.co.jp/weather/jp/earthquake/list/"
	listRes, err := client.Get(URL)
	if err != nil {
		return nil, err
	}
	defer listRes.Body.Close()

	if listRes.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("在获取日本地震信息列表的过程中返回状态码异常: %d", listRes.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(listRes.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
