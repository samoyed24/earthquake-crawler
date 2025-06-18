package app

import (
	"earthquake-crawler/internal/config"
	"earthquake-crawler/internal/task"
	"time"

	"github.com/sirupsen/logrus"
)

func runTask(taskFunc func(), interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		go taskFunc()
		<-ticker.C
	}
}

func RunApp() error {
	logrus.Info("爬虫程序开始运行")
	if config.Cfg.CrawlerSwitch.JapanEarthquakeCrawlerSwitch {
		logrus.Infof("[日本地震信息]已添加爬虫任务，间隔%v秒执行", config.Cfg.CrawlerInterval.JapanEarthquakeInterval)
		go runTask(task.JapanEarthquakeCrawlTask, time.Duration(config.Cfg.CrawlerInterval.JapanEarthquakeInterval)*time.Second)
	}
	if config.Cfg.CrawlerSwitch.JapanEEWCrawlerSwitch {
		logrus.Infof("[日本EEW]已添加爬虫任务，间隔%v秒执行", config.Cfg.CrawlerInterval.JapanEEWInterval)
		go runTask(task.JapanEEWCrawlTask, time.Duration(config.Cfg.CrawlerInterval.JapanEEWInterval)*time.Second)

	}
	select {}
}
