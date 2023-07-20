package environment

import (
	"edoex/environment/flags"
	"edoex/utils/sliceutils"
	"fmt"
)

func GetSimulator() (string, error) {
	if flags.Simulator != "" {
		if ValidateSimulator(flags.Simulator) {
			return flags.Simulator, nil
		}

		return "", fmt.Errorf("Simulator '%s' is not available", flags.Simulator)
	}

	if Config.Simulator != "" {
		if ValidateSimulator(Config.Simulator) {
			return Config.Simulator, nil
		}

		return "", fmt.Errorf("Simulator '%s' is not available", Config.Simulator)
	}

	return "", fmt.Errorf("A simulator must be stated through the flag '--simulator' or 'edoex.config.yaml'")
}

func ValidateSimulator(simulator string) bool {
	return sliceutils.Contains(availableSimulators, simulator)
}

const (
	OmegaSimulator  = "omega"
	EdoproSimulator = "edopro"
)

var availableSimulators []string = []string{
	EdoproSimulator,
	OmegaSimulator,
}
