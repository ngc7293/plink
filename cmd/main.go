package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "plink",
	Short: "CLI tools for PrusaLink",
	RunE: func(cmd *cobra.Command, args []string) error {
		// The default behaviour is to show the current job status
		return ShowStatus()
	},
}

func init() {
	rootCmd.AddCommand(storageCmd)
	storageCmd.AddCommand(lsCmd)
	storageCmd.AddCommand(rmCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
