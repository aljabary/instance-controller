/*
Copyright © 2022 Ibnu Nugraha <nugrahaberbakti@gmail.com>

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

// start represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Use this command for start instance",
	Long:  `Run instance by start, use instance ID`,
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
			fmt.Println("You must supply an instance ID (-i INSTANCE-ID")
			return
		}
		var dr = false

		input := &ec2.StartInstancesInput{
			InstanceIds: []string{
				instanceID,
			},
			DryRun: &dr,
		}

		_, err = types.StartInstance(context.TODO(), client, input)
		if err != nil {
			fmt.Println("Got an error starting the instance")
			fmt.Println(err)
			return
		}

		fmt.Println("Started instance with ID " + instanceID)

	},
}

func init() {

	rootCmd.AddCommand(startCmd)
}
