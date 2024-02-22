package wendover

import (
	"os"

	"github.com/spf13/cobra"

	_ "github.com/derhabicht/wendover/internal/config"
	"github.com/derhabicht/wendover/internal/logging"
)

var logLevel string

var rootCmd = &cobra.Command{
	Use:   "wendover",
	Short: "Helper tooling for Civil Air Patrol activities",
	Long:  ``,
}

func Execute(version string) {
	rootCmd.Version = version
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&logLevel, "loglevel", "info", "")

	logging.InitLogging(logLevel, true)
}
