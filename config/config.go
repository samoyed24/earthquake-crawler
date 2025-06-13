package config

type HttpRequestConfig struct {
	TimeoutSeconds int `toml:"timeout_seconds"`
}

type DBConfig struct {
	DBPath string `toml:"db_path"`
}

type CrawlerSwitchConfig struct {
	JapanEarthquakeCrawlerSwitch bool `toml:"japan_earthquake_crawler_switch"`
}

type CrawlerIntervalConfig struct {
	JapanEarthquakeInterval int `toml:"japan_earthquake_interval"`
}

type Config struct {
	HttpRequest     HttpRequestConfig     `toml:"httpRequest"`
	DB              DBConfig              `toml:"db"`
	CrawlerSwitch   CrawlerSwitchConfig   `toml:"crawlerSwitch"`
	CrawlerInterval CrawlerIntervalConfig `toml:"crawlerInterval"`
}

var Cfg Config
