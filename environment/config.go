package environment

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type config struct {
	ExpansionName string `yaml:"expansion"`
}

var Config config

func init() {
	configFile, err := os.ReadFile(ConfigPath())
	if err != nil {
		log.Fatalf("Cannot read '%s'\n", ConfigPath())
	}

	yaml.Unmarshal(configFile, &Config)
}
