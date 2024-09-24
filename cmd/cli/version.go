package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of SailHost",
	Long:  `All software has versions. This is SailHost's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("SailHost CLI version 0.0.1")
	},
}
