package config

import "github.com/spf13/cobra"

type ConfigOption struct {
	// Path of config file
	Path string
	// Cobra cmd
	Cmd *cobra.Command
}

type Option func(o *ConfigOption)

func WithConfigPath(path string) Option {
	return func(o *ConfigOption) {
		o.Path = path
	}
}
