package email

import (
	"earthquake-crawler/internal/config"
	"fmt"

	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

var emailConfig = &config.Cfg.Email

func SendEmail(toList []string, subject string, content string) error {
	// logrus.Info("1")
	m := gomail.NewMessage()
	m.SetHeader("From", emailConfig.Username)
	m.SetHeader("From", "alias"+"<"+emailConfig.Username+">")
	m.SetHeader("Bcc", toList...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)

	d := gomail.NewDialer(
		emailConfig.Host,
		emailConfig.Port,
		emailConfig.Username,
		emailConfig.Password,
	)
	for i := 1; i <= emailConfig.MaxRetries; i++ {
		if err := d.DialAndSend(m); err != nil {
			logrus.Errorf("发送邮件失败(第%d次尝试), 错误原因: %v", i, err)
		} else {
			return nil
		}
	}
	return fmt.Errorf("发送邮件失败(已尝试%d次), 请检查配置与网络", emailConfig.MaxRetries)
}
