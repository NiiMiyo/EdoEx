package environment

import (
	"edoex/logger"
	"os"

	"gopkg.in/yaml.v3"
)

type config struct {
	ExpansionName string `yaml:"expansion"`
	EdoproPath    string `yaml:"edopro_path"`
	OmegaPath     string `yaml:"omega_path"`
	Simulator     string `yaml:"simulator"`
}

var Config config

func UpdateConfig() {
	configFile, err := os.ReadFile(SourceConfigPath())
	if err != nil {
		logger.ErrorfErr("Cannot read '%s'", err, SourceConfigPath())
	}

	yaml.Unmarshal(configFile, &Config)

	var globalConfig config
	globalConfigFile, err := os.ReadFile(GlobalConfigPath())
	if err != nil {
		logger.ErrorfErr("Cannot read '%s'", err, GlobalConfigPath())
	}

	yaml.Unmarshal(globalConfigFile, &globalConfig)
	if Config.EdoproPath == "" {
		Config.EdoproPath = globalConfig.EdoproPath
	}

	if Config.OmegaPath == "" {
		Config.OmegaPath = globalConfig.OmegaPath
	}

	if Config.Simulator == "" {
		Config.Simulator = globalConfig.Simulator
	}
}
