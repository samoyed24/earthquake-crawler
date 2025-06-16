package jpquakecrawler

import (
	"earthquake-crawler/internal/config"
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func GetJapanEarthquakeListDoc() (*goquery.Document, error) {
	client := &http.Client{
		Timeout: time.Duration(config.Cfg.HttpRequest.TimeoutSeconds) * time.Second,
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
