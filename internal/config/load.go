package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

func LoadConfig() error {
	Path := "data/config.toml"
	f, err := os.Open(Path)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := toml.NewDecoder(f).Decode(&Cfg); err != nil {
		return fmt.Errorf("配置文件解析失败: %v", err)
	}
	return nil
}
