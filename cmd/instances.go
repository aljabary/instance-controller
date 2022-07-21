/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	types "github/aljabary/instance-controller/types/ConfigFile"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
)

// instancesCmd represents the instances command
var instancesCmd = &cobra.Command{
	Use:   "instances",
	Short: "Register instances to controlled",
	Long: `Register intances to contrroled.
	use: instances [intanceid1, instanceid2]`,
	Run: func(cmd *cobra.Command, args []string) {
		data, err := ioutil.ReadFile("configinstances.json")
		if err != nil {
			log.Fatal("Failed read config file")
		}
		configfile := &types.ConfigFile{}
		json.Unmarshal(data, configfile)
		if len(args) > 0 {
			configfile.Instances = args
			b, _ := json.Marshal(configfile)
			ioutil.WriteFile("configinstances.json", b, 0777)
		} else {
			for _, id := range configfile.Instances {
				log.Print(id)
			}
		}
		/*

			data, err := ioutil.ReadFile("configinstances.json")
			if err != nil {
				log.Info("Failed read config file")
			}
			js := json.Unmarshal(data,)
			mydata := []byte("All the data I wish to write to a file")
			ioutil.WriteFile("configinstances.json", mydata, 0777)
		*/
	},
}

func init() {
	rootCmd.AddCommand(instancesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// instancesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// instancesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
