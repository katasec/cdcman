/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/katasec/cdcman/ddl"
	"github.com/spf13/cobra"
)

// ddlCmd represents the ddl command
var ddlCmd = &cobra.Command{
	Use:   "ddl",
	Short: "Opertions to extract and apply the DDL for your database tables",
	Long:  `Opertions to extract and apply the DDL for your database tables.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("ddl called")
		ddl.Init()
	},
}

func init() {
	rootCmd.AddCommand(ddlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ddlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ddlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
