# Overview
EC2Cycle is a simple terminal tool to start, stop and view the state of an 
AWS EC2 instance written in golang

EC2Cycle provides:
* Easy to use and simple terminal tool
* Creates empty config.yaml file if not provided
* Config requires minimum variables needed to establish AWS session

# Concepts
* Use of [viper](https://github.com/spf13/viper) for configuration manager.
* Use of [cobra](https://github.com/spf13/cobra) for CLI command scaffolding
* Use of [AWS SDK for GO](https://aws.amazon.com/sdk-for-go/) libraries to create a AWS session and AWS EC2 Client instance 

# Commands
* View the state of instance `./ec2cycle state`
* Start the instance `./ec2cycle start`
* Stop the instance `./ec2cycle stop`

# Config
The `config.yaml` file needs to exist in the same directory as the executable `./ec2cycle` binary

The following entries need to exist in the config file:
```yaml
accesskey: <your aws access key>
secretaccesskey: <aws secret access key>
ec2instance: <aws ec2 instance id>
region: ap-southeast-2
```
# Whats next
* Better error handling
* Specify custom config file as argument
* Add Buildfile
* Add documentation

# Changelog
## Version 0.1alpha
Initial, very basic version