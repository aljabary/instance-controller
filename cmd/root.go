/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "instance-controller",
	Short: "Schedule your running instances",
	Long: `
	 ___________________________
	|                           |
	| Instance Controller v.1.0 |
	| Author: Ibnu Nugraha      |
	|___________________________|
	
	
	chedule when your instance runs and when to stop. 
	Scheduling instances automatically saves costs.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//Run: func(cmd *cobra.Command, args []string) { fmt.Println("halo") },
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.instance-controller.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().StringP("instance-id", "i", "", "Instance ID")
	rootCmd.PersistentFlags().StringP("timestart", "t", "", "cronjob time format for start schedule")
	rootCmd.PersistentFlags().StringP("timestop", "s", "", "cronjob time format for stop schedule")
	rootCmd.PersistentFlags().StringP("region", "r", "", "AWS Region")
	rootCmd.PersistentFlags().StringP("profile", "p", "default", "AWS Profile, let for default")
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
