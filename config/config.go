package config

import (
	"bufio"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type AppConfig interface {
	GetAccessKey() *string
	GetAWSRegion() *string
	GetEC2Instance() *string
	GetSecretAccessKey() *string
}

type fileConfig struct {
	accessKey       string
	secretAccessKey string
	region          string
	ec2Instance     string
}

func (ac *fileConfig) GetAccessKey() *string {
	return &ac.accessKey
}

func (ac *fileConfig) GetAWSRegion() *string {
	return &ac.region
}

func (ac *fileConfig) GetEC2Instance() *string {
	return &ac.ec2Instance
}

func (ac *fileConfig) GetSecretAccessKey() *string {
	return &ac.secretAccessKey
}

func LoadConfig() AppConfig {
	ac := fileConfig{}
	apc := newConfigFile()
	vp := viper.New()
	vp.SetConfigType("yaml")
	if !apc.exists() {
		initializeConfig(apc, vp)
	}
	file, err := os.Open(apc.path())
	if err != nil {
		fmt.Printf("Unexpected error while reading config %v", err)
	}
	err = vp.ReadConfig(bufio.NewReader(file))
	if err != nil {
		fmt.Printf("Unexpected error while reading config %v", err)
	}
	ac.region = fmt.Sprint(vp.Get("region"))
	ac.accessKey = fmt.Sprint(vp.Get("accessKey"))
	ac.ec2Instance = fmt.Sprint(vp.Get("ec2Instance"))
	ac.secretAccessKey = fmt.Sprint(vp.Get("secretAccessKey"))
	return &ac
}

func initializeConfig(apc configFile, vp *viper.Viper) {
	err := apc.create()
	if err != nil {
		fmt.Printf("Could not create config file [%v]\n", apc.create())
	}
	vp.Set("region", "ap-southeast-2")
	vp.Set("accessKey", "<access key>")
	vp.Set("ec2Instance", "<instance id>")
	vp.Set("secretAccessKey", "<access secret key>")
	err = vp.WriteConfigAs(apc.path())
	if err != nil {
		fmt.Printf("Serious error %v \n", err)
	}
	fmt.Println("Config file did not exit. Empty config file was created.\nPlease configure and run app again")
	os.Exit(1)
}