/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"github/aljabary/instance-controller/types"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Use this command for stop insance",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		region, _ := rootCmd.Flags().GetString("region")
		if region == "" {
			fmt.Println("You must supply an region name (-r regnion_name)")
			return
		}
		profile, _ := rootCmd.Flags().GetString("profile")
		if profile == "" {
			fmt.Println("You must supply an profile name (-p profile_name)")
			return
		}
		cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region), config.WithSharedConfigProfile(profile))
		instanceID, _ := rootCmd.Flags().GetString("instance-id")
		client := ec2.NewFromConfig(cfg)
		if instanceID == "" {
			fmt.Println("You must supply an instance ID (-i INSTANCE-ID)")
			return
		}
		fmt.Println("Try to stop instance " + instanceID)
		var dr = false
		input := &ec2.StopInstancesInput{
			InstanceIds: []string{
				instanceID,
			},
			DryRun: &dr,
		}

		_, err = types.StopInstance(context.TODO(), client, input)
		if err != nil {
			fmt.Println("Got an error stoping the instance")
			fmt.Println(err)
			return
		}

		fmt.Println("Stoped instance with ID " + instanceID)

	},
}

func init() {

	rootCmd.AddCommand(stopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
