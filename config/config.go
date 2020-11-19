/**
 * Author: kekai wang
 */
package config

import (
	"io/ioutil"
	"log"
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

type Config struct {
}

var (
	Env        = os.Getenv("env")
	configPath = map[string]string{
		"dev":  "/config/",
		"test": "/config/",
		"prod": "/config/",
	}
)

var (
	config Config
	once   sync.Once
)

func Get() Config {
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

	return config
}
