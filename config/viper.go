package config

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type viperConfig struct {
	opts  ConfigOption
	viper *viper.Viper
}

var viperOnce sync.Once

func (v *viperConfig) Init(opts ...Option) error {
	for _, o := range opts {
		o(&v.opts)
	}
	return v.Load()
}

func (v *viperConfig) Get(key string) interface{} {
	return v.viper.Get(key)
}

func (v *viperConfig) Scan(key string, to interface{}) error {
	return v.viper.UnmarshalKey(key, to)
}

func (v *viperConfig) Set(key string, value interface{}) {
	v.viper.Set(key, value)
}

func (v *viperConfig) Load() error {
	if v.opts.Path != "" {
		// Use config file from the flag.
		v.viper.SetConfigFile(v.opts.Path)
	} else {
		// Find config directory.
		cfgDir, err := os.UserConfigDir()
		if err != nil {
			return err
		}

		// Search config in config directory with name ".logent" (without extension).
		v.viper.AddConfigPath(cfgDir)
		v.viper.SetConfigName(".log-aggregator")
	}

	v.viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	err := v.viper.ReadInConfig()
	if err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
	v.bindFlags()
	return err
}

func (v *viperConfig) String() string {
	return "Viper config"
}

func (v *viperConfig) bindFlags() {
	if v.opts.Cmd == nil {
		return
	}
	v.opts.Cmd.Flags().VisitAll(func(f *pflag.Flag) {
		// Environment variables can't have dashes in them, so bind them to their equivalent
		// keys with underscores, e.g. --favorite-color to STING_FAVORITE_COLOR
		if strings.Contains(f.Name, "-") {
			envVarSuffix := strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_"))
			v.viper.BindEnv(f.Name, envVarSuffix)
		}

		// Apply the viper config value to the flag when the flag is not set and viper has a value
		if !f.Changed && v.viper.IsSet(f.Name) {
			val := v.Get(f.Name)
			v.opts.Cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})
}

func NewViperConfig(opts ...Option) (Config, error) {
	c := &viperConfig{
		viper: viper.New(),
	}
	err := c.Init(opts...)
	if err != nil {
		return nil, err
	}
	viperOnce.Do(func() {
		defaultConfig = c
	})
	return c, nil

}

func WithCobraCmd(cmd *cobra.Command) Option {
	return func(o *ConfigOption) {
		o.Cmd = cmd
	}
}
