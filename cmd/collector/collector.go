package collector

import (
	"log"

	"github.com/itzmanish/go-logent/pkg/config"
	"github.com/spf13/cobra"
)

func RunCollector(cmd *cobra.Command, args []string) {
	watchers := config.Watchers{}
	err := config.Scan("watchers", &watchers)
	if err != nil {
		log.Fatal(err)
	}
}
