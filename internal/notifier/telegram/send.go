package telegram

import (
	"earthquake-crawler/internal/config"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"sync"
)

func sendSingleMessage(bot *tgbotapi.BotAPI, content string, receiver int64) error {
	msg := tgbotapi.NewMessage(receiver, content)
	for i := 1; i <= config.Cfg.Telegram.MaxRetries; i++ {
		_, err := bot.Send(msg)
		if err != nil {
			logrus.Errorf("[Notifier-Telegram]Telegram Bot信息发送失败(已尝试%d次): %v", i, err)
		} else {
			return nil
		}
	}
	return fmt.Errorf("[Notifier-邮件]发送邮件失败(已尝试%d次), 请检查配置与网络", config.Cfg.Telegram.MaxRetries)
}

func SendTelegramMessage(content string, receivers []int64) error {
	bot, err := tgbotapi.NewBotAPI(config.Cfg.Telegram.BotToken)
	if err != nil {
		return fmt.Errorf("[Notifier-Telegram]在初始化Telegram bot的过程中发生错误: %v", err)
	}

	sem := make(chan struct{}, config.Cfg.Telegram.MaxSendOnceATime)
	var wg sync.WaitGroup

	for _, receiver := range receivers {
		sem <- struct{}{}
		wg.Add(1)

		go func(r int64) {
			defer wg.Done()
			defer func() { <-sem }()

			err := sendSingleMessage(bot, content, r)
			if err != nil {
				logrus.Errorf("[Notifier-Telegram]在向%v发送信息时发生错误: %v", r, err)
			}
		}(receiver)
	}

	wg.Wait()
	return nil
}
