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
	Enable           bool `toml:"enable"`
	CrawlInterval    int  `toml:"crawl_interval"`
	ParseAfterMinute int  `toml:"parse_after_minute"`
}

type JPEEWConfig struct {
	Enable        bool   `toml:"enable"`
	CrawlInterval int    `toml:"crawl_interval"`
	RedisEnable   bool   `toml:"redis_enable"`
	RedisKey      string `toml:"redis_key"`
}

type EmailReceiveJPQuakeConfig struct {
	Receive        bool `toml:"receive"`
	MaxReceiveOnce int  `toml:"max_receive_once"`
}

type EmailReceiveJPEEWConfig struct {
	Receive          bool `toml:"receive"`
	ReceiveAlertOnly bool `toml:"receive_alert_only"`
	ReceiveTrain     bool `toml:"receive_train"`
}

type EmailReceiveConfig struct {
	ReceiverUsers       []string                  `toml:"receive_users"`
	EmailReceiveJPQuake EmailReceiveJPQuakeConfig `toml:"jpquake"`
	EmailReceiveJPEEW   EmailReceiveJPEEWConfig   `toml:"jpeew"`
}

type EmailConfig struct {
	Enable       bool               `toml:"enable"`
	Host         string             `toml:"host"`
	Port         int                `toml:"port"`
	Username     string             `toml:"username"`
	Password     string             `toml:"password"`
	MaxRetries   int                `toml:"max_retries"`
	EmailReceive EmailReceiveConfig `toml:"receive"`
}

type Config struct {
	HttpRequest HttpRequestConfig `toml:"httpRequest"`
	Params      ParamsConfig      `toml:"params"`
	Redis       RedisConfig       `toml:"redis"`
	DB          DBConfig          `toml:"db"`
	JPQuake     JPQuakeConfig     `toml:"jpquake"`
	JPEEW       JPEEWConfig       `toml:"jpeew"`
	Email       EmailConfig       `toml:"email"`
}

var Cfg Config
