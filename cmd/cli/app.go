package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(appCmd)
	rootCmd.AddCommand(initAppCmd)
}

var appCmd = &cobra.Command{
	Use:   "app",
	Short: "Run SailHost web service",
	Long:  `Run SailHost web service`,
	Run: func(cmd *cobra.Command, args []string) {
		if !isRoot() {
			fmt.Println("You must run SailHost as root user. Please run 'sudo sailhost app'")
			return
		}

		// TODO: Run SailHost web service
	},
}

var initAppCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize SailHost application",
	Long:  `Initialize SailHost application`,
	Run: func(cmd *cobra.Command, args []string) {
		if !isRoot() {
			fmt.Println("You must run SailHost as root user. Please run 'sudo sailhost init'")
			return
		}

		// TODO: Initialize SailHost application
	},
}
