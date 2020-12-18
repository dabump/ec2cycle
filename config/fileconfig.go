package config

import (
	"bufio"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func LoadConfig() (AppConfig, error) {
	ac := config{}
	apc := newConfigFile()
	vp := viper.New()
	vp.SetConfigType("yaml")
	if !apc.exists() {
		initializeConfig(apc, vp)
	}
	file, err := os.Open(apc.path())
	if err != nil {
		return nil, fmt.Errorf("unexpected error while opening config file\n%v", err)
	}
	err = vp.ReadConfig(bufio.NewReader(file))
	if err != nil {
		return nil, fmt.Errorf("unexpected error while reading config file\n%v", err)
	}
	if region := vp.Get("region"); region != nil {
		ac.region = fmt.Sprint(region)
	}
	if accessKey := vp.Get("accessKey"); accessKey != nil {
		ac.accessKey = fmt.Sprint(accessKey)
	}
	if ec2Instance := vp.Get("ec2Instance"); ec2Instance != nil {
		ac.ec2Instance = fmt.Sprint(ec2Instance)
	}
	if secretAccessKey := vp.Get("secretAccessKey"); secretAccessKey != nil {
		ac.secretAccessKey = fmt.Sprint(secretAccessKey)
	}
	err = ac.Valid()
	if err != nil {
		return nil, fmt.Errorf("missing config\n%v", err)
	}
	return &ac, nil
}

func initializeConfig(apc configFile, vp *viper.Viper) {
	err := apc.create()
	if err != nil {
		fmt.Printf("Could not create config file [%v]\n", apc.create())
		os.Exit(1)
	}
	vp.Set("region", "ap-southeast-2")
	vp.Set("accessKey", "<access key>")
	vp.Set("ec2Instance", "<instance id>")
	vp.Set("secretAccessKey", "<access secret key>")
	err = vp.WriteConfigAs(apc.path())
	if err != nil {
		fmt.Printf("unexpected error during write attempt of config file\n%v\n", err)
		os.Exit(1)
	}
	fmt.Println("Config file did not exit. Empty config file was created.\nPlease configure and run app again")
	os.Exit(0)
}