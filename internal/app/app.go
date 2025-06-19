package app

import (
	"earthquake-crawler/internal/config"
	"earthquake-crawler/internal/crawler/task"
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
	logrus.Info("[聚合地震信息爬虫]程序开始运行")
	if config.Cfg.JPQuake.Enable {
		logrus.Infof("[日本地震信息]已添加爬虫任务，间隔%v秒执行", config.Cfg.JPQuake.CrawlInterval)
		go runTask(task.JapanEarthquakeCrawlTask, time.Duration(config.Cfg.JPQuake.CrawlInterval)*time.Second)
	}
	if config.Cfg.JPEEW.Enable {
		logrus.Infof("[日本EEW]已添加爬虫任务，间隔%v秒执行", config.Cfg.JPEEW.CrawlInterval)
		go runTask(task.JapanEEWCrawlTask, time.Duration(config.Cfg.JPQuake.CrawlInterval)*time.Second)
	}
	select {}
}
