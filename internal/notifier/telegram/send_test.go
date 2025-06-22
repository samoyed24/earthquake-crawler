package telegram

import (
	"earthquake-crawler/internal/config"
	"testing"
)

func TestSendTelegramMessage(t *testing.T) {
	config.Cfg.Telegram.Enable = true
	config.Cfg.Telegram.BotToken = ""
	config.Cfg.Telegram.Receive.ReceiverUsers = []int64{}
	config.Cfg.Telegram.MaxRetries = 3
	err := SendTelegramMessage("Test Message", config.Cfg.Telegram.Receive.ReceiverUsers)
	if err != nil {
		t.Error(err)
	}
}
