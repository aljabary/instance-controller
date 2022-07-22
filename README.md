# instance-controller
CLI for set schdule instance run and stop, this helps you to saving cost. For example, when you run dev staging server (EC2), that doesn't need 24 hours for running.
The staging or dev server just need to run for 8-9 hours/day, just for work time in week days. So you need set run and stop the server, if you do manually it will takes time.
So you need automatically, this is the right tools for you, you can set the schedule for automation.

## Requirement
- [AWS CLI v.2](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-welcome.html)

## Config
Make sure you have configure AWS credentials with AWS SDK CLI, following:

```golang
aws configure 
```
You must provide AWS Access Key and AWS Access Secret

## Getting Started
Following step by step bellow:
### 1. download binary file
You must download the binary file ```instance-controller``` and file ```configinstances.json``` from ```build``` folder and put on your server controller. So in this case you should have 1 server as controller, usually a devOps always have at least 1 server for controller other resouce service and other servers.
#### OR
you can download from release with curl:
```golang
curl -L https://github.com/aljabary/instance-controller/releases/download/beta/instance-controller.zip instance-controller.zip
```
and then extract to folder:
```golang
unzip instance-controller.zip -d ./myfolder
```
### 2. set AWS region
```golang
cd myfolder 
./instance-controller jobs awsregion -r "ap-southeast-1"
```
replace ```ap-southeast-1``` with your region
### 3. set AWS profile
```golang
./instance-controller jobs awsprofilen -p default
```
replace ```default``` with your credential profile
### 4. Register instances
You must register server instances one or more with instance ID
```golang
./instance-controller instances i-abcxxxxx i-defxxxxx i-ghixxxx
```
to read list of instances:
```golang
./instance-controller instances 
```
### 5. set Schedule
To automatic run and stop you must set the schedule for ```t``` (run) and ```s``` (stop)
```golang
./instance-controller jobs schedule -t "0 0 7 * * *" -s "0 0 18 * * *" 
```
This example we set sechdule for run every 7 am and stop instances every 18 pm for every day. time format such as cronjob, ```second minute clock day month year```.

### 6. Running
this will running jobs automaticly for run and stop instances
```golang
./instance-controller jobs run
```
