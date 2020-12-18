package main

import (
	"ec2cycle/aws"
	"ec2cycle/config"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var version = "v0.2alpha"

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
		ac, err := config.LoadConfig()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		awsSession := aws.NewSession(ac)
		state, err := awsSession.State(ac.GetEC2Instance())
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Printf("Instance state: %s\n", state)
	},
}

var instanceStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start instance",
	Long:  `Starts the AWS instance`,
	Run: func(cmd *cobra.Command, args []string) {
		ac, err := config.LoadConfig()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		awsSession := aws.NewSession(ac)
		awsSession.Start(ac.GetEC2Instance())
		fmt.Println("Starting instance")
	},
}

var instanceStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop instance",
	Long:  `Stops the AWS instance`,
	Run: func(cmd *cobra.Command, args []string) {
		ac, err := config.LoadConfig()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		awsSession := aws.NewSession(ac)
		awsSession.Stop(ac.GetEC2Instance())
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
