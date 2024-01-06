/*
Copyright © 2024 Aravindh

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)
var version = "v0.0.1 a1"
// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Gives you the version of the extrak program",
	Long: `Version command gives you the extrak program version number for example v1.0`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
