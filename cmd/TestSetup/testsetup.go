/*
Copyright © 2024 Aravindh

*/
package TestSetup

import (
	"fmt"

	"github.com/spf13/cobra"
)

// TestsetupCmd represents the testsetup command
var TestsetupCmd = &cobra.Command{
	Use:   "testsetup",
	Short: "Extrak setup test command",
	Long: `This is a test setup command to test your extrak application setup`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("testsetup called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testsetupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testsetupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
