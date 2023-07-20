package environment

import (
	"path/filepath"
)

const (
	OmegaImagesDir  = "Arts"
	OmegaScriptsDir = "Scripts"
)

func OmegaDatabaseBuildPath() string {
	return filepath.Join(BuildPath(), Config.ExpansionName+".db")
}

func OmegaImagesBuildPath() string {
	return filepath.Join(BuildPath(), OmegaImagesDir)
}

func OmegaScriptsBuildPath() string {
	return filepath.Join(BuildPath(), OmegaScriptsDir)
}

func (self *config) OmegaSyncPath() string {
	return filepath.Join(self.OmegaPath, "YGO Omega_Data/Files")
}
