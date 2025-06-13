package app

import (
	"japan-earthquake-webspider/config"
	"japan-earthquake-webspider/internal/task"
	"time"

	"github.com/sirupsen/logrus"
)

func runTask(taskFunc func() error, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		taskFunc()
		<-ticker.C
	}
}

func RunApp() error {
	logrus.Info("爬虫程序开始运行")
	if config.Cfg.CrawlerSwitch.JapanEarthquakeCrawlerSwitch {
		logrus.Infof("已添加日本地震信息爬虫任务，间隔%v秒执行", config.Cfg.CrawlerInterval.JapanEarthquakeInterval)
		go runTask(task.JapanEarthquakeCrawlTask, time.Duration(config.Cfg.CrawlerInterval.JapanEarthquakeInterval)*time.Second)
	}
	select {}
}
