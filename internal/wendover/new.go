package wendover

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/derhabicht/wendover/internal/activity"
	"github.com/derhabicht/wendover/internal/logging"
)

var newCmd = &cobra.Command{
	Use:   "new [ActivityName]",
	Short: "Create a new CAP activity",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run:   runNewCmd,
}

func runNewCmd(cmd *cobra.Command, args []string) {
	logger := logging.Logger{}

	err := activity.CreateNewActivity(args[0], logger)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to create new activity.")
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(newCmd)
}
