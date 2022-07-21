/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github/aljabary/instance-controller/cmd"
	"log"

	"github.com/robfig/cron"
)

func printCronEntries(cronEntries []*cron.Entry) {
	log.Print("Cron Info: %+v\n", cronEntries)
}
func cJob() {
	log.Print("Create new cron")
	c := cron.New()
	c.AddFunc("*/5 * * * *", func() { log.Print("") })

	// Start cron with one scheduled job
	log.Print("Start cron")
	c.Start()
	printCronEntries(c.Entries())
}

func main() {

	/*sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan,
		syscall.SIGINT,
		syscall.SIGKILL,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	<-sigchan
	*/
	cmd.Execute()
	//time.Sleep(1 * time.Minute)
}
