package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	AppName    string `json:"app_name"`
	AppVersion string `json:"app_version"`
	DbDriver   string `json:"db_driver"`
	DbName     string `json:"db_name"`
	DbUser     string `json:"db_user"`
	DbPass     string `json:"db_pass"`
	DbHost     string `json:"db_host"`
	DbPort     int    `json:"db_port"`
}

func Set(conf Config) {
	res, _ := json.Marshal(conf)
	err := ioutil.WriteFile("config/config.json", res, 0777)

	if err != nil {
		fmt.Println("Error when write config.json :", err)
	}
}

func Get() {
	res, err := ioutil.ReadFile("config/config.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(res))
}
