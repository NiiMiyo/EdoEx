package environment

import (
	"os"
	"path/filepath"
)

var (
	WorkingDir, _ = filepath.Abs(".")
	ProgramDir    = filepath.Dir(os.Args[0])
)

const (
	BuildDir           = "build"
	SourceMetaDir      = "meta"
	SourceCardsDir     = "cards"
	SourceScriptsDir   = "scripts"
	GlobalTemplatesDir = "card_templates"
	BuildPicsDir       = "pics"
	SourceArtworksDir  = "artworks"
	BuildScriptDir     = "script"

	StringsFile = "strings.conf"
	ConfigFile  = "edoex.config.yaml"
)

func BuildPath() string {
	path, _ := filepath.Abs(filepath.Join(WorkingDir, BuildDir))
	return path
}

func SourceMetaPath() string {
	path, _ := filepath.Abs(filepath.Join(WorkingDir, SourceMetaDir))
	return path
}

func SourceCardsPath() string {
	path, _ := filepath.Abs(filepath.Join(WorkingDir, SourceCardsDir))
	return path
}

func BuildStringsPath() string {
	return filepath.Join(BuildPath(), StringsFile)
}

func BuildDatabasePath() string {
	return filepath.Join(BuildPath(), Config.ExpansionName+".cdb")
}

func SourceConfigPath() string {
	return filepath.Join(WorkingDir, ConfigFile)
}

func GlobalConfigPath() string {
	return filepath.Join(ProgramDir, ConfigFile)
}

func SourceScriptsPath() string {
	return filepath.Join(WorkingDir, SourceScriptsDir)
}

func GlobalTemplatesPath() string {
	return filepath.Join(ProgramDir, GlobalTemplatesDir)
}

func BuildPicsPath() string {
	return filepath.Join(BuildPath(), BuildPicsDir)
}

func SourceArtworksPath() string {
	return filepath.Join(WorkingDir, SourceArtworksDir)
}

func BuildScriptsPath() string {
	return filepath.Join(BuildPath(), BuildScriptDir)
}
