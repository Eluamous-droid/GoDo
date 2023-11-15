/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package command

import (
	"os"
	"strconv"

	"github.com/eluamous-droid/godo/pkg/tools"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "Todo application that syncs tasks between machines",
	Long: `Todo app written in go, is synced with the api server if that is running`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		for i, item := range tools.ReadTodosFromFile().TodoItems {
			println(strconv.Itoa(i) + ": Group: " + item.Group + " Task: " + item.Task)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
