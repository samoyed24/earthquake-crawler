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

func GetJapanEEW(queryTime string) (*model.RawJapanEEWData, error) {
	// 这里需要保证下个请求开始前上一个请求一定结束，如果没有结束就放弃。
	// 建议不要把配置文件中的超时秒数设得太小（默认1s），不然会永远都拿不到数据。
	eewTimeout := time.Duration(min(config.Cfg.HttpRequest.TimeoutSeconds))
	client := &http.Client{
		Timeout: eewTimeout * time.Second,
	}
	// queryTime = "20240101161430"
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

	var data model.RawJapanEEWData
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
