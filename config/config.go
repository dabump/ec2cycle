package config

import "errors"

type AppConfig interface {
	GetAccessKey() *string
	GetAWSRegion() *string
	GetEC2Instance() *string
	GetSecretAccessKey() *string
}

type config struct {
	accessKey       string
	secretAccessKey string
	region          string
	ec2Instance     string
}

func (ac *config) GetAccessKey() *string {
	return &ac.accessKey
}

func (ac *config) GetAWSRegion() *string {
	return &ac.region
}

func (ac *config) GetEC2Instance() *string {
	return &ac.ec2Instance
}

func (ac *config) GetSecretAccessKey() *string {
	return &ac.secretAccessKey
}

func (ac *config) Valid() error {
	if ac.accessKey == "" {
		return errors.New("accessKey value was not specified in the config file")
	}
	if ac.secretAccessKey == "" {
		return errors.New("secretAccessKey value was not specified in the config file")
	}
	if ac.region == "" {
		return errors.New("region value was not specified in the config file")
	}
	if ac.ec2Instance == "" {
		return errors.New("ec2Instance value was not specified in the config file")
	}
	return nil
}
