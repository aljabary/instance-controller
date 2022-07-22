/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"

	instance "github/aljabary/instance-controller/types"
	types "github/aljabary/instance-controller/types/ConfigFile"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/robfig/cron"
	"github.com/spf13/cobra"
)

var ConfigFile types.ConfigFile

func cJob(configfile *types.ConfigFile) {
	c := cron.New()
	c.AddFunc(configfile.ScheduleStart, func() {
		if len(configfile.Instances) > 0 {
			cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(configfile.Awsregion), config.WithSharedConfigProfile(configfile.Awsprofile))
			for _, instanceID := range configfile.Instances {
				fmt.Println("Starting instance " + instanceID)
				client := ec2.NewFromConfig(cfg)
				var dr = false
				input := &ec2.StartInstancesInput{
					InstanceIds: []string{
						instanceID,
					},
					DryRun: &dr,
				}
				_, err = instance.StartInstance(context.TODO(), client, input)
				if err != nil {
					fmt.Println("Got an error starting the instance " + instanceID)
					fmt.Println(err)
					return
				}
			}
		}
	})
	c.Start()
}

func cJobStop(configfile *types.ConfigFile) {
	c := cron.New()
	c.AddFunc(configfile.ScheduleStop, func() {
		if len(configfile.Instances) > 0 {
			cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(configfile.Awsregion), config.WithSharedConfigProfile(configfile.Awsprofile))
			for _, instanceID := range configfile.Instances {
				fmt.Println("Stopping instance " + instanceID)
				client := ec2.NewFromConfig(cfg)
				var dr = false
				input := &ec2.StopInstancesInput{
					InstanceIds: []string{
						instanceID,
					},
					DryRun: &dr,
				}
				_, err = instance.StopInstance(context.TODO(), client, input)
				if err != nil {
					fmt.Println("Got an error starting the instance " + instanceID)
					fmt.Println(err)
					return
				}
			}

		}

	})
	c.Start()
}

var x int

// jobsCmd represents the jobs command
var jobsCmd = &cobra.Command{
	Use:   "jobs",
	Short: "Run or Stop CronJob",
	Long:  `Run or Stop CronJob schedulling`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("need one argument")
		}

		data, err := ioutil.ReadFile("configinstances.json")
		if err != nil {
			log.Fatal("Failed read config file")
		}
		configfile := &types.ConfigFile{}
		json.Unmarshal(data, configfile)
		if args[0] == "run" {
			log.Print("running jobs")
			sigchan := make(chan os.Signal, 1)
			cJob(configfile)
			cJobStop(configfile)
			signal.Notify(sigchan,
				syscall.SIGINT,
				syscall.SIGKILL,
				syscall.SIGTERM,
				syscall.SIGQUIT)
			<-sigchan

		}
		if args[0] == "stop" {
		}
		if args[0] == "schedule" {
			schedule, _ := rootCmd.Flags().GetString("timestart")
			stop, _ := rootCmd.Flags().GetString("timestop")
			if schedule == "" && stop == "" {
				fmt.Print("please provide time for cronjob")
				return
			}
			configfile.ScheduleStart = schedule
			configfile.ScheduleStop = stop
			b, _ := json.Marshal(configfile)
			ioutil.WriteFile("configinstances.json", b, 0777)
			fmt.Print("Schedule has setup" + schedule)
		}
		if args[0] == "awsregion" {
			region, _ := rootCmd.Flags().GetString("region")
			if region == "" {
				fmt.Println("You must supply a region name (-r regnion_name)")
				return
			}
			configfile.Awsregion = region
			b, _ := json.Marshal(configfile)
			ioutil.WriteFile("configinstances.json", b, 0777)
			fmt.Print("Aws region has setup" + region)
		}
		if args[0] == "awsprofile" {
			profile, _ := rootCmd.Flags().GetString("profile")
			if profile == "" {
				fmt.Println("You must supply a profile name (-p profile_name)")
				return
			}
			configfile.Awsregion = profile
			b, _ := json.Marshal(configfile)
			ioutil.WriteFile("configinstances.json", b, 0777)
			fmt.Print("AWS Profile has setup" + profile)

		}
	},
}

func init() {
	jobsCmd.SetHelpFunc(func(cmd *cobra.Command, str []string) {
		fmt.Println(cmd.Use + "  " + cmd.Short)
		fmt.Println("usage:")
		fmt.Println("	jobs run 				: run jobs scheduling")
		//fmt.Println("	jobs stop 				: stop jobs scheduling")
		fmt.Println("	jobs schedule -t [start time cronjob format] -s  [top time cronjob format]	: set cronjob schedule time, ex: jobs schedule -t '*/30 * * * *'")
		fmt.Println("	jobs awsregion -r [region-name]		: set aws region")
		fmt.Println("	jobs awsprofile -p [profilename]	: set aws profile credential")
	})
	rootCmd.AddCommand(jobsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jobsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// jobsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
