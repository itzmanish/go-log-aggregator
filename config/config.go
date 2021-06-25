package config

type Config interface {
	Init(opts ...Option) error
	Get(key string) interface{}
	Set(key string, value interface{})
	Scan(key string, value interface{}) error
	Load() error
	String() string
}

type Watcher struct {
	Watch string                   `mapstructure:"watch" json:"watch"`
	Tags  []map[string]interface{} `mapstructure:"tags" json:"tags"`
}

type Watchers []Watcher

type ServerConfig struct {
	Host               string `mapstructure:"host"`
	Port               string `mapstructure:"port"`
	Store              string `mapstructure:"store"`
	Bucket             string `mapstructure:"bucket"`
	Endpoint           string `mapstructure:"endpoint"`
	AwsAccessKey       string `mapstructure:"aws_access_key"`
	AwsSecretKey       string `mapstructure:"aws_secret_key"`
	PathAddressingMode string `mapstructure:"path_addressing_mode"`
}

type AgentConfig struct {
	Watchers `mapstructure:"watchers"`
	Server   ServerConfig      `mapstructure:"server"`
	Retry    map[string]string `mapstructure:"retry"`
}

var defaultConfig Config

func Init(opts ...Option) error {
	return defaultConfig.Init(opts...)
}

func Get(key string) interface{} {
	return defaultConfig.Get(key)
}

func Scan(key string, to interface{}) error {
	return defaultConfig.Scan(key, to)
}

func Set(key string, value interface{}) {
	defaultConfig.Set(key, value)
}

func Load() error {
	return defaultConfig.Load()
}

func String() string {
	return defaultConfig.String()
}

func NewConfig(opts ...Option) (Config, error) {
	return NewViperConfig(opts...)
}
