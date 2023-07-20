package environment

import (
	"path/filepath"
)

const (
	EdoproImagesDir   = "pics"
	EdoproScriptsDir  = "script"
	EdoproStringsFile = "strings.conf"
)

func EdoproStringsBuildPath() string {
	return filepath.Join(BuildPath(), EdoproStringsFile)
}

func EdoproDatabaseBuildPath() string {
	return filepath.Join(BuildPath(), Config.ExpansionName+".cdb")
}

func EdoproImagesBuildPath() string {
	return filepath.Join(BuildPath(), EdoproImagesDir)
}

func EdoproScriptsBuildPath() string {
	return filepath.Join(BuildPath(), EdoproScriptsDir)
}

func (self *config) EdoproSyncZipPath() string {
	return filepath.Join(self.EdoproPath, "expansions", self.ExpansionName+".zip")
}
