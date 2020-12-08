package main

import (
	"ec2cycle/internal/config"
	"ec2cycle/internal/ec2remote"
	"fmt"
    "github.com/spf13/cobra"
	"os"
)

var version = "v0.1alpha"

var rootCmd = &cobra.Command{
	Use:   "ec2cycle",
	Short: "Over simplistic EC2 power cycle",
	Long: `ec2cycle allows the user to start, stop and view state of a single
ec2 instance. Clean and simple.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var instanceStateCmd = &cobra.Command{
	Use:   "state",
	Short: "Output instance state",
	Long:  `Pokes at AWS to fetch the state of the AWS instance`,
	Run: func(cmd *cobra.Command, args []string) {
		ac := config.LoadConfig()
		awsSession := ec2remote.NewSession(ac)
		state, err := awsSession.InstanceState(ac.GetEC2Instance())
		if err != nil {
			fmt.Printf("Error %v", err)
		}
		fmt.Printf("Instance state: %s\n", state)	},
}

var instanceStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start instance",
	Long:  `Starts the AWS instance`,
	Run: func(cmd *cobra.Command, args []string) {
		ac := config.LoadConfig()
		awsSession := ec2remote.NewSession(ac)
		awsSession.StartInstance(ac.GetEC2Instance())
		fmt.Println("Starting instance")
	},
}

var instanceStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop instance",
	Long:  `Stops the AWS instance`,
	Run: func(cmd *cobra.Command, args []string) {
		ac := config.LoadConfig()
		awsSession := ec2remote.NewSession(ac)
		awsSession.StopInstance(ac.GetEC2Instance())
		fmt.Println("Stopping instance")
	},
}

func main() {
	fmt.Printf("Welcome to EC2 cycle %s\n", version)
	rootCmd.AddCommand(instanceStateCmd)
	rootCmd.AddCommand(instanceStartCmd)
	rootCmd.AddCommand(instanceStopCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}