/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"log"

	"os/exec"
	"syscall"

	"github.com/eiannone/keyboard"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	ini "gopkg.in/ini.v1"
)

// pcCmd represents the pc command
var pcCmd = &cobra.Command{
	Use:   "pc",
	Short: "aws profile change",
	Long:  `A command help your change your aws profile in ~/.aws/config, will export AWS_PROFILE in your env.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pc called")

		var defaultAwsConfigPath = os.Getenv("HOME") + "/.aws/config"
		cfg, err := ini.Load(defaultAwsConfigPath)
		var items = []string{}
		for _, profile := range cfg.Sections() {
			if strings.Contains(profile.Name(), "profile") {
				items = append(items, strings.Split(profile.Name(), " ")[1])
			}
		}
		if err != nil {
			log.Fatalf("Fail to read file: %v", color.RedString("%s", err))
			os.Exit(1)
		}

		selectedIndex := 0

		err = clearTerminal()
		if err != nil {
			log.Fatal("Can not clean Terminal...:", err)
		}

		printMenu(items, selectedIndex)

		err = keyboard.Open()
		if err != nil {
			log.Fatal("Can not open keyboard...:", err)
		}
		defer func() {
			_ = keyboard.Close()
		}()

		for {
			char, key, err := keyboard.GetKey()
			if err != nil {
				log.Fatal("Can not get key:", err)
			}

			if key == keyboard.KeyArrowUp {
				selectedIndex = (selectedIndex - 1 + len(items)) % len(items)
			} else if key == keyboard.KeyArrowDown {
				selectedIndex = (selectedIndex + 1) % len(items)
			} else if key == keyboard.KeyEnter {
				break
			}

			err = clearTerminal()
			if err != nil {
				log.Fatal("Can not clean Terminal...:", err)
			}

			printMenu(items, selectedIndex)

			if char == 'q' || key == keyboard.KeyCtrlC {
				fmt.Println("User press Ctrl+C.")
				os.Exit(0)
			}
		}

		fmt.Printf("You choose: %s\n", items[selectedIndex])
		file, err := os.Create("/tmp/shared_env.txt")
		if err != nil {
			fmt.Println("Error creating shared file:", err)
			return
		}
		defer file.Close()

		_, err = fmt.Fprintf(file, "AWS_PROFILE=%s\n", items[selectedIndex])
		if err != nil {
			fmt.Println("Error writing to shared file:", err)
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(pcCmd)
}

func printMenu(items []string, selectedIndex int) {
	for i, item := range items {
		if i == selectedIndex {
			fmt.Printf("%s", color.BlueString("-> %s\n", item))
		} else {
			fmt.Printf("   %s\n", item)
		}
	}
	fmt.Println("Use 'Enter' to choose, use 'q' bye")
}

func clearTerminal() error {
	cmdClear := exec.Command("clear") // for Linux and MacOS
	if syscall.Environ()[0] == "windows" {
		cmdClear = exec.Command("cmd", "/c", "cls") // for Windows
	}
	cmdClear.Stdout = os.Stdout
	return cmdClear.Run()
}
