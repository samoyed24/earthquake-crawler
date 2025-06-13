package spider

import (
	"fmt"
	"japan-earthquake-webspider/internal/config"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func GetEarthquakeDetailDoc(eqTime string) (*goquery.Document, error) {
	client := &http.Client{
		Timeout: time.Duration(config.Cfg.HttpRequest.TimeoutSeconds) * time.Second,
	}
	URL := fmt.Sprintf("https://typhoon.yahoo.co.jp/weather/jp/earthquake/%v.html", eqTime)
	resp, err := client.Get(URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("在获取地震详情信息的过程中返回状态码异常: %v", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}
