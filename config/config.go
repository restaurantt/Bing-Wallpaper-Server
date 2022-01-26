/**
 * @Author: Ne-21
 * @Description:
 * @File:  config.go
 * @Version: 1.0.0
 * @Date: 2022/1/26 13:29
 */

package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	MySql *MySql
	Redis *Redis
}

type MySql struct {
	Address  string `yaml:"address"`
	Port     string `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	UserName string `yaml:"username"`
	PassWord string `yaml:"password"`
}

type Redis struct {
	Address     string `yaml:"address"`
	Port        string `yaml:"port"`
	MaxIdle     int    `yaml:"maxIdle"`
	MaxActive   int    `yaml:"maxActive"`
	IdleTimeout int    `yaml:"idleTimeout"`
}

var AppConfig *Config

func InitConfig() {
	AppConfig = &Config{}
	content, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("解析config.yaml读取错误: %v", err)
	}

	fmt.Println(string(content))
	if yaml.Unmarshal(content, &AppConfig) != nil {
		log.Fatalf("解析config.yaml出错: %v", err)
	}
}
