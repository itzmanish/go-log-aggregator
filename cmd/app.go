/*
Copyright © 2021 Manish <itzmanish108@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/itzmanish/go-loganalyzer/config"
	"github.com/spf13/cobra"
)

var cfgFile string

// appCmd represents the base command when called without any subcommands
var appCmd = &cobra.Command{
	Use:     "loganalyzer",
	Short:   "loganalyzer is Log analyzer agent which runs in host system and sends log from logfiles",
	Version: "1.0.0",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the appCmd.
func Execute() {
	cobra.CheckErr(appCmd.Execute())
}

func init() {
	cobra.OnInitialize(func() {
		_, err := config.NewViperConfig(config.WithConfigPath(cfgFile))
		cobra.CheckErr(err)
	})

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	appCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/.loganalyzer.yaml)")

	// // Cobra also supports local flags, which will only run
	// // when this action is called directly.
	// appCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}