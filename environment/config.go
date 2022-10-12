package environment

import (
	"edoex/logger"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type config struct {
	ExpansionName string `yaml:"expansion"`
	Gamedir       string `yaml:"gamedir"`
}

var Config config

func UpdateConfig() {
	configFile, err := os.ReadFile(SourceConfigPath())
	if err != nil {
		logger.ErrorfErr("Cannot read '%s'", err, SourceConfigPath())
	}

	yaml.Unmarshal(configFile, &Config)

	if Config.Gamedir == "" {
		var globalConfig config

		globalConfigFile, err := os.ReadFile(GlobalConfigPath())
		if err != nil {
			logger.ErrorfErr("Cannot read '%s'", err, GlobalConfigPath())
		}

		yaml.Unmarshal(globalConfigFile, &globalConfig)

		Config.Gamedir = globalConfig.Gamedir
	}
}

func (self *config) ExpansionSyncPath() string {
	return filepath.Join(self.Gamedir, "expansions", self.ExpansionName+".zip")
}
