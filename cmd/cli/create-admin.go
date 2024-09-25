package cli

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(createAdminCmd)
}

var createAdminCmd = &cobra.Command{
	Use:   "create-admin",
	Short: "Create a new admin user",
	Long:  `Create a new admin user`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Create a new admin user
	},
}
