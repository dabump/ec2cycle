package config

import (
	"bufio"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type AppConfig struct {
	accessKey       string
	secretAccessKey string
	region          string
	ec2Instance     string
}

func (ac *AppConfig) GetAccessKey() *string {
	return &ac.accessKey
}

func (ac *AppConfig) GetSecretAccessKey() *string {
	return &ac.secretAccessKey
}

func (ac *AppConfig) GetAWSRegion() *string {
	return &ac.region
}

func (ac *AppConfig) GetEC2Instance() *string {
	return &ac.ec2Instance
}

func LoadConfig() *AppConfig {
	ac := AppConfig{}
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
	ac.accessKey = fmt.Sprint(vp.Get("accessKey"))
	ac.secretAccessKey = fmt.Sprint(vp.Get("secretAccessKey"))
	ac.ec2Instance = fmt.Sprint(vp.Get("ec2Instance"))
	ac.region = fmt.Sprint(vp.Get("region"))
	return &ac
}

func initializeConfig(apc *configFile, vp *viper.Viper) {
	err := apc.create()
	if err != nil {
		fmt.Printf("Could not create config file [%v]\n", apc.create())
	}
	vp.Set("accessKey", "<access key>")
	vp.Set("secretAccessKey", "<access secret key>")
	vp.Set("ec2Instance", "<instance id>")
	vp.Set("region", "ap-southeast-2")
	err = vp.WriteConfigAs(apc.path())
	if err != nil {
		fmt.Printf("Serious error %v \n", err)
	}
	fmt.Println("Config file did not exit. Empty config file was created.\nPlease configure and run app again")
	os.Exit(1)
}
