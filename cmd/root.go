/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/goggle-source/gowiki/internal/commands"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gowiki",
	Short: "statistical website generator on go",
	Long:  `This is my pet project, in which I have implemented reading md files, metadata of these files, as well as reading html templates. Using the contents of md files and templates, I create a ready-made html page that is served by the http server. The project is configured using YML config`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gowiki.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(commands.InitBuild())
	rootCmd.AddCommand(commands.InitServe())
	rootCmd.AddCommand(commands.InitNew())
	rootCmd.AddCommand(commands.InitCreateEnv())
}
