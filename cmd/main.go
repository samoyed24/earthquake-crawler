package main

import (
	"earthquake-crawler/internal/app"
	"earthquake-crawler/internal/config"
	"earthquake-crawler/internal/storage"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("在读取配置的过程中发生错误：%v", err)
	}
	log.Info("配置读取成功")
	if err := storage.LoadDB(); err != nil {
		log.Fatalf("在读取SQLite数据库的过程中发生错误: %v", err)
	}
	log.Info("SQLite数据库读取成功")
	if config.Cfg.Redis.Enable {
		if err := storage.InitRedisClient(); err != nil {
			log.Fatalf("连接Redis失败: %v", err)
		}
		log.Info("Redis连接成功")
	}
	runAppErr := make(chan error, 1)
	go func() {
		runAppErr <- app.RunApp()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-quit:
		log.Info("程序正在退出...")
	case err := <-runAppErr:
		if err != nil {
			log.Errorf("爬虫程序启动失败: %v", err)
		}
	}
	if err := storage.DB.Close(); err != nil {
		log.Errorf("SQLite数据库连接关闭失败: %v", err)
	} else {
		log.Info("SQLite数据库连接已关闭")
	}
	if config.Cfg.Redis.Enable {
		if err := storage.CloseRedisClient(); err != nil {
			log.Errorf("Redis连接关闭失败: %v", err)
		} else {
			log.Info("Redis连接已关闭")
		}
	}
	log.Info("程序退出")
}
