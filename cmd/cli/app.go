package cli

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(appCmd)
	appCmd.AddCommand(initAppCmd)
}

var appCmd = &cobra.Command{
	Use:   "app",
	Short: "Run SailHost web service",
	Long:  `Run SailHost web service`,
	Run: func(cmd *cobra.Command, args []string) {
		//
	},
}

var initAppCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize SailHost application",
	Long:  `Initialize SailHost application`,
	Run: func(cmd *cobra.Command, args []string) {
		//
	},
}
