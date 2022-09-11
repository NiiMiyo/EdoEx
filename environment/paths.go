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
	BuildDir     = "build"
	MetaDir      = "meta"
	CardsDir     = "cards"
	ScriptsDir   = "scripts"
	ImagesDir    = "images"
	TemplatesDir = "card_templates"
	PicsDir      = "pics"

	StringsFile = "strings.conf"
	ConfigFile  = "edoex.config.yaml"
)

func BuildPath() string {
	path, _ := filepath.Abs(filepath.Join(WorkingDir, BuildDir))
	return path
}

func MetaPath() string {
	path, _ := filepath.Abs(filepath.Join(WorkingDir, MetaDir))
	return path
}

func CardsPath() string {
	path, _ := filepath.Abs(filepath.Join(WorkingDir, CardsDir))
	return path
}

func StringsPath() string {
	return filepath.Join(BuildPath(), StringsFile)
}

func DatabasePath() string {
	return filepath.Join(BuildPath(), Config.ExpansionName+".cdb")
}

func ConfigPath() string {
	return filepath.Join(WorkingDir, ConfigFile)
}

func GlobalConfigPath() string {
	return filepath.Join(ProgramDir, ConfigFile)
}

func ScriptsPath() string {
	return filepath.Join(WorkingDir, ScriptsDir)
}

func ImagesPath() string {
	return filepath.Join(WorkingDir, ImagesDir)
}

func TemplatesPath() string {
	return filepath.Join(ProgramDir, TemplatesDir)
}

func PicsPath() string {
	return filepath.Join(BuildPath(), PicsDir)
}
