package main

import (
	"japan-earthquake-webspider/config"
	"japan-earthquake-webspider/internal/app"
	"japan-earthquake-webspider/internal/storage"

	log "github.com/sirupsen/logrus"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("在读取配置的过程中发生错误：%v", err)
	}
	log.Info("配置读取成功")
	if err := storage.LoadDB(); err != nil {
		log.Fatalf("在读取数据库的过程中发生错误：%v", err)
	}
	log.Info("数据库读取成功")
	if err := app.RunEarthquakeListSpider(); err != nil {
		log.Fatalf("爬虫程序运行失败：%v", err)
	}
	// log.Info("爬虫程序开始运行")
}
