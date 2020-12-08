package ec2remote

import (
	"ec2cycle/internal/config"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"os"
)

type AWSSession struct {
	awsSession *session.Session
	ec2Client  *ec2.EC2
}

func NewSession(config *config.AppConfig) *AWSSession {
	ses := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(*config.GetAWSRegion()),
		Credentials: credentials.NewStaticCredentials(*config.GetAccessKey(), *config.GetSecretAccessKey(), ""),
	}))
	ec2Client := ec2.New(ses)
	return &AWSSession{awsSession: ses, ec2Client: ec2Client}
}

func (as *AWSSession) StartInstance(instanceID *string) {
	var ids []*string
	ids = append(ids, instanceID)
	input := ec2.StartInstancesInput{InstanceIds: ids}
	_, err := as.ec2Client.StartInstances(&input)
	if err != nil {
		fmt.Printf("Error during ec2 interaction %v", err)
		os.Exit(2)
	}
}

func (as *AWSSession) StopInstance(instanceID *string) {
	var ids []*string
	ids = append(ids, instanceID)
	input := ec2.StopInstancesInput{InstanceIds: ids}
	_, err := as.ec2Client.StopInstances(&input)
	if err != nil {
		fmt.Printf("Error during ec2 interaction %v", err)
		os.Exit(2)
	}
}

func (as *AWSSession) InstanceState(instanceID *string) (string, error) {
	var ids []*string
	ids = append(ids, instanceID)
	allInstance := true
	input := ec2.DescribeInstanceStatusInput{
		InstanceIds:         ids,
		IncludeAllInstances: &allInstance,
	}
	output, err := as.ec2Client.DescribeInstanceStatus(&input)
	if err != nil {
		fmt.Printf("Error during ec2 interaction %v", err)
		os.Exit(2)
	}
	for _, status := range output.InstanceStatuses {
		if *status.InstanceId == *instanceID {
			return *status.InstanceState.Name, nil
		}
	}
	return "", errors.New("could not determine instance state")
}
