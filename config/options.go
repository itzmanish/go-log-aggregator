package config

type ConfigOption struct {
	// Path of config file
	Path string
}

type Option func(o *ConfigOption)

func WithConfigPath(path string) Option {
	return func(o *ConfigOption) {
		o.Path = path
	}
}
