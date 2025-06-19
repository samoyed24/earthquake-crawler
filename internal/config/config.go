package config

type HttpRequestConfig struct {
	TimeoutSeconds int `toml:"timeout_seconds"`
}

type ParamsConfig struct {
	Timezone string `toml:"timezone"`
}

type DBConfig struct {
	DBPath string `toml:"db_path"`
}

type RedisConfig struct {
	Enable   bool   `toml:"enable"`
	Addr     string `toml:"addr"`
	Port     int    `toml:"port"`
	Password string `toml:"password"`
	DB       int    `toml:"db"`
}

type JPQuakeConfig struct {
	Enable        bool `toml:"enable"`
	CrawlInterval int  `toml:"crawl_interval"`
}

type JPEEWConfig struct {
	Enable        bool   `toml:"enable"`
	CrawlInterval int    `toml:"crawl_interval"`
	RedisEnable   bool   `toml:"redis_enable"`
	RedisKey      string `toml:"redis_key"`
}

type Config struct {
	HttpRequest HttpRequestConfig `toml:"httpRequest"`
	Params      ParamsConfig      `toml:"params"`
	Redis       RedisConfig       `toml:"redis"`
	DB          DBConfig          `toml:"db"`
	JPQuake     JPQuakeConfig     `toml:"jpquake"`
	JPEEW       JPEEWConfig       `toml:"jpeew"`
}

var Cfg Config
