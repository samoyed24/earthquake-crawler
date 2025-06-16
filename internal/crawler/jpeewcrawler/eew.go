package jpeewcrawler

import (
	"earthquake-crawler/internal/config"
	"earthquake-crawler/internal/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func GetJapanEEW(queryTime string) (*model.JapanEEWData, error) {
	client := &http.Client{
		Timeout: time.Duration(config.Cfg.HttpRequest.TimeoutSeconds) * time.Second,
	}
	URL := fmt.Sprintf("http://www.kmoni.bosai.go.jp/webservice/hypo/eew/%v.json", queryTime)
	res, err := client.Get(URL)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var data model.JapanEEWData
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
