/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"time"

	"log"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	ini "gopkg.in/ini.v1"
)

// tokenCmd represents the token command
var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "A check aws expiration time",
	Long:  `Get the aws expiration of AWS_PROFILE`,
	Run: func(cmd *cobra.Command, _ []string) {

		var defaultFilePath = os.Getenv("HOME") + "/.aws/credentials"
		profile, err := cmd.Flags().GetString("profile")
		if err != nil {
			log.Fatalf("Input string: %v", color.RedString("%s", err))
			os.Exit(1)
		}
		filePath, _ := cmd.Flags().GetString("file")
		var cfg *ini.File
		if filePath == "default" {
			cfg, err = ini.Load(defaultFilePath)
			filePath = defaultFilePath
		} else {
			cfg, err = ini.Load(filePath)
		}

		if profile == "AWS_PROFILE" {
			profile = os.Getenv("AWS_PROFILE")
		}

		if err != nil {
			log.Fatalf("Fail to read file: %v", color.RedString("%s", err))
			os.Exit(1)
		}
		log.Printf("Try find profile_name: %v in the %v ...\n", color.GreenString("%s", profile), color.GreenString("%s", filePath))
		aws_expiration := cfg.Section(profile).Key("aws_expiration").String()
		if aws_expiration == "" {
			log.Fatalf("Can not find %v profile in the %v", color.RedString("%s", profile), color.RedString("%s", filePath))
			os.Exit(1)
		}
		now, err := time.Parse(time.RFC3339, cfg.Section(profile).Key("aws_expiration").String())
		if err != nil {
			log.Fatalln("Parse time failed: ", color.RedString("%s", err))
			os.Exit(1)
		}
		asiaTaipei, _ := time.LoadLocation("Asia/Taipei")
		origin := time.Date(
			now.Year(),
			now.Month(),
			now.Day(),
			now.Hour(),
			now.Minute(),
			now.Second(),
			now.Nanosecond(), time.UTC)
		color.Blue("AWS Expiration Date: %s", origin.In(asiaTaipei))

		color.Red("Time left: %s", time.Until(origin.In(asiaTaipei)))
	},
}

func init() {
	rootCmd.AddCommand(tokenCmd)
	tokenCmd.PersistentFlags().String("profile", "AWS_PROFILE", "$AWS_PROFILE env")
	tokenCmd.Flags().String("file", "default", "The file of credentials")
}
