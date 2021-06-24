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
	"os"
	"os/signal"

	"github.com/itzmanish/go-loganalyzer/handler"
	"github.com/itzmanish/go-loganalyzer/internal/logger"
	"github.com/itzmanish/go-loganalyzer/internal/server"
	"github.com/itzmanish/go-loganalyzer/internal/store"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Log analyzer server to collect logs from agent and process it.",
	Run: func(cmd *cobra.Command, args []string) {
		port, err := cmd.Flags().GetString("port")
		if err != nil {
			logger.Fatal(err)
		}
		storeName, err := cmd.Flags().GetString("store")
		if err != nil {
			logger.Fatal(err)
		}
		sProvider, ok := store.Stores[storeName]
		if !ok {
			logger.Fatal("Store Provider doesn't exits. please select a different store provider.")
		}
		sopts := []store.Option{}
		if bn, err := cmd.Flags().GetString("bucket"); bn != "" && err == nil {
			sopts = append(sopts, store.WithDirectory(bn))
		}
		if ep, err := cmd.Flags().GetString("endpoint"); ep != "" && err == nil {
			sopts = append(sopts, store.WithS3Endpoint(ep))
		}
		if epm, err := cmd.Flags().GetBool("pathaddressingmode"); epm && err == nil {
			sopts = append(sopts, store.WithPathStyleAddressing(epm))
		}
		if ak, err := cmd.Flags().GetString("aws_access_key"); ak != "" && err == nil {
			sopts = append(sopts, store.WithAWSAccessKey(ak))
		}
		if sak, err := cmd.Flags().GetString("aws_secret_key"); sak != "" && err == nil {
			sopts = append(sopts, store.WithAWSSecretAccessKey(sak))
		}

		store, err := sProvider(sopts...)
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info("Selected storage backend is ", store.String())
		hdl := handler.NewHandler(store)
		s := server.NewServer(server.WithPort(port), server.WithHandler(hdl))
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt)
		go func() {
			<-exit
			if err := s.Stop(); err != nil {
				logger.Fatal(err)
			}
		}()
		if err := s.Start(); err != nil {
			logger.Error(err)
		}
	},
}

func init() {
	appCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	serverCmd.Flags().StringP("port", "p", "33555", "log analyzer server port")
	serverCmd.Flags().StringP("store", "s", "file", "Store backend for logs")
	serverCmd.Flags().String("endpoint", "", "S3 enpoint")
	serverCmd.Flags().Bool("pathaddressingmode", true, `S3 differentiates between "virtual hosted bucket addressing" and "path-style addressing".
Example URL for the virtual host style: http://BUCKET.s3.amazonaws.com/KEY.
Example UL for the path style: http://s3.amazonaws.com/BUCKET/KEY.
Most S3-compatible servers work with both styles,
but some work only with the virtual host style 
(e.g. Alibaba Cloud Object Storage Service (OSS))
and some work only with the path style 
(especially self-hosted services like a Minio server running on localhost).`)
	serverCmd.Flags().String("bucket", "", "Bucket Name for s3 and Directory name for local fs")
	serverCmd.Flags().String("aws_access_key", "", `AWS Access key for s3.
Optional (read from shared credentials file or environment variable if not set).
Environment variable: 'AWS_ACCESS_KEY_ID'.`)
	serverCmd.Flags().String("aws_secret_key", "", `AWS Secret key for s3.
Optional (read from shared credentials file or environment variable if not set).
Environment variable: "AWS_SECRET_ACCESS_KEY".`)
}
