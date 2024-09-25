package cli

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(restartCmd)
}

var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart SailHost Web Service",
	Long:  `Restart SailHost Web Service`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Restart SailHost
	},
}
