package email

import (
	"earthquake-crawler/internal/config"
	"testing"
)

func TestSendEmail(t *testing.T) {
	// 测试邮件是否能够正常发送，请在这里将关于邮箱的参数填入后测试是否能正常发送邮件
	// 测试通过后再写到config.toml中
	config.Cfg.Email.Host = "example@example.com"
	config.Cfg.Email.Port = -1
	config.Cfg.Email.MaxRetries = 3
	config.Cfg.Email.Username = ""
	config.Cfg.Email.Password = ""
	config.Cfg.Email.EmailReceive.ReceiverUsers = []string{"example@example.com"}
	err := SendEmail(config.Cfg.Email.EmailReceive.ReceiverUsers, "Test Email", "<h1>Hello World</h1>")
	if err != nil {
		t.Fatalf("[测试]邮件发送测试未通过: %v", err)
	}
}
