/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package TestSetup

import (
	"fmt"

	"github.com/spf13/cobra"
)

var databaseType string;
// testdatabaseCmd represents the testdatabase command
var testdatabaseCmd = &cobra.Command{
	Use:   "testdatabase",
	Short: "Test your database connectivity",
	Long: `This function is used to test the database connectivity`,
	Run: func(cmd *cobra.Command, args []string) {
    if(databaseType == "1"){
      fmt.Println("sqllite")
    }else if(databaseType == "2"){
      fmt.Println("postgres")
    }else{
      fmt.Println("Unsupported databaseType")
    }
	},
}

func init() {
  testdatabaseCmd.Flags().StringVarP(&databaseType, "databaseType", "d", "", "Database type 1.sqllite, 2.postgres")
  TestsetupCmd.AddCommand(testdatabaseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testdatabaseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testdatabaseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
