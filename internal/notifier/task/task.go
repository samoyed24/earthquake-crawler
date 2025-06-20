package task

import (
	"earthquake-crawler/internal/config"
	"earthquake-crawler/internal/model"
	"earthquake-crawler/internal/notifier/email"
	"fmt"

	"github.com/sirupsen/logrus"
)

func SendJPQuakeEmail(data *model.JapanEarthquakeDetail) {
	logrus.Info("[Notifier-邮件]正在尝试向已设置的邮件接受者发送日本地震信息")
	content, err := email.RenderJapanEarthquakeEmailTemplate(data)
	if err != nil {
		logrus.Errorf("[Notifier-邮件]在渲染日本地震信息的过程中发生错误: %v", err)
		return
	}
	subject := fmt.Sprintf("[日本地震情报更新]%v发生地震", data.Center)
	err = email.SendEmail(config.Cfg.Email.EmailReceive.ReceiverUsers, subject, *content)
	if err != nil {
		logrus.Errorf("[Notifier-邮件]在发送日本地震信息邮件的过程中发生错误: %v", err)
	}
	logrus.Info("[Notifier-邮件]邮件发送成功")
}
