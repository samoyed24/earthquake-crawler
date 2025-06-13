package config

type HttpRequestConfig struct {
	TimeoutSeconds int `toml:"timeout_seconds"`
}

type DBConfig struct {
	DBPath string `toml:"db_path"`
}

type Config struct {
	HttpRequest HttpRequestConfig `toml:"httpRequest"`
	DB          DBConfig          `toml:"db"`
}

var Cfg Config
