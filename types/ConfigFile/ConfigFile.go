/*
Type Configfile interface
*/
package types

type ConfigFile struct {
	Instances     []string
	ScheduleStart string
	ScheduleStop  string
	Awsregion     string
	Awsprofile    string
}
