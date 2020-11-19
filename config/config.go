/**
 * Author: kekai wang
 */
package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

type Config struct {
	App struct {
		Env  string `yaml:"env"`
		Name string `yaml:"name"`
	} `yaml:"app"`

	Mysql struct {
		Uri             string `yaml:"uri"`
		Host            string `yaml:"host"`
		User            string `yaml:"user"`
		Pass            string `yaml:"pass"`
		Port            string `yaml:"port"`
		DbName          string `yaml:"db_name"`
		MaxOpenConns    int    `yaml:"max_open_conns"`
		MaxIdleConns    int    `yaml:"max_idle_conns"`
		ConnMaxLifeTime int    `yaml:"conn_max_lifetime"`
		LogMode         bool   `yaml:"log_mode"`
	} `yaml:"mysql"`
}

var (
	Env        = os.Getenv("env")
	configPath = map[string]string{
		"dev":  "config/",
		"test": "config/",
		"prod": "config/",
	}
)

var (
	config Config
	once   sync.Once
)

func Get() Config {
	fmt.Println("-----------", os.Getenv("env"))
	once.Do(func() {
		configFile := configPath[Env] + "config." + Env + ".yaml"

		configData, err := ioutil.ReadFile(configFile)
		if err != nil {
			log.Printf("Load config err %s", err)
		}

		err = yaml.Unmarshal([]byte(configData), &config)
		if err != nil {
			log.Printf("Yaml unmarshal err: %s", err)
		}
	})
	fmt.Println("-----------", os.Getenv("env"), config)
	return config
}
