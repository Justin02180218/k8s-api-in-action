package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/viper"
)

func InitConfig() {
	err := initConfigByStream()
	if err != nil {
		fmt.Println("Read local config file.")
		initConfigByLocal()
	}
}

func initConfigByLocal() {
	workDir, _ := os.Getwd()
	viper.AddConfigPath(workDir + "/config")
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initConfigByStream() (err error) {
	values, err := ioutil.ReadFile("etc/config/application.yaml")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	viper.SetConfigType("yaml")
	err = viper.ReadConfig(bytes.NewReader(values))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}
