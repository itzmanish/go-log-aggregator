package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/spf13/viper"
)

type viperConfig struct {
	Opts ConfigOption
}

var viperOnce sync.Once

func (v *viperConfig) Init(opts ...Option) error {
	for _, o := range opts {
		o(&v.Opts)
	}
	return v.Load()
}

func (v *viperConfig) Get(key string) interface{} {
	return viper.Get(key)
}

func (v *viperConfig) Scan(key string, to interface{}) error {
	return viper.UnmarshalKey(key, to)
}

func (v *viperConfig) Set(key string, value interface{}) {
	viper.Set(key, value)
}

func (v *viperConfig) Load() error {
	if v.Opts.Path != "" {
		// Use config file from the flag.
		viper.SetConfigFile(v.Opts.Path)
	} else {
		// Find config directory.
		cfgDir, err := os.UserConfigDir()
		if err != nil {
			return err
		}

		// Search config in config directory with name ".logent" (without extension).
		viper.AddConfigPath(cfgDir)
		viper.SetConfigName(".logent")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
	return err
}

func (v *viperConfig) String() string {
	return "Viper config"
}

func NewViperConfig(opts ...Option) (Config, error) {
	c := &viperConfig{}
	err := c.Init(opts...)
	if err != nil {
		return nil, err
	}
	viperOnce.Do(func() {
		defaultConfig = c
	})
	return c, nil

}
