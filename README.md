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
