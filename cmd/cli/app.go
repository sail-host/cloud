package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var devMode bool

func init() {
	appCmd.Flags().BoolVarP(&devMode, "dev", "d", false, "Run SailHost in development mode")
	rootCmd.AddCommand(appCmd)
}

var appCmd = &cobra.Command{
	Use:   "app",
	Short: "Run SailHost web service",
	Long:  `Run SailHost web service`,
	Run: func(cmd *cobra.Command, args []string) {
		if devMode && !isRoot() {
			fmt.Println("You must run SailHost as root user. Please run 'sudo sailhost app'")
			return
		}

		if devMode && !isLinux() {
			fmt.Println("SailHost is only supported on Linux")
			return
		}

		// TODO: Run SailHost web service
	},
}
