package config

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

//go:embed config.template.toml
var configTemplate embed.FS

func LoadConfig() error {
	// template := "data/config.template.toml"
	path := "data/config.toml"
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, 0755)
	f, err := os.Open(path)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		logrus.Info("未找到配置文件, 正在尝试创建新的配置文件")
		// err := util.CopyFile(template, path)
		content, err := configTemplate.ReadFile("config.template.toml")
		if err != nil {
			return fmt.Errorf("无法读取配置模板: %v", err)
		}
		err = os.WriteFile(path, content, 0644)
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
