/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/neilkuan/aws-token-exp/pkg/constants"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version",
	Long:  `version.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(constants.NAME, "version", constants.VERSION)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
