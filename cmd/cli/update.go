package cli

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update SailHost",
	Long:  `Update SailHost`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Update SailHost
	},
}
