package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sailhost",
	Short: "SailHost is a cloud hosting platform",
	Long: `A Fast and Flexible Cloud Hosting Platform built with
love by SailHost in Go.
Complete documentation is available at https://sailhost.io`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
