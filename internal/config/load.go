package config

import (
	"earthquake-crawler/internal/util"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

func LoadConfig() error {
	template := "data/config.template.toml"
	path := "data/config.toml"
	f, err := os.Open(path)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		logrus.Info("未找到配置文件, 正在尝试创建新的配置文件")
		err := util.CopyFile(template, path)
		if err != nil {
			return fmt.Errorf("创建新的配置文件失败: %v", err)
		}
		logrus.Infof("已创建新的配置文件于%v, 请根据需要自行配置", path)
		f, err = os.Open(path)
		if err != nil {
			return err
		}
	}
	defer f.Close()

	if _, err := toml.NewDecoder(f).Decode(&Cfg); err != nil {
		return fmt.Errorf("配置文件解析失败: %v", err)
	}
	return nil
}
