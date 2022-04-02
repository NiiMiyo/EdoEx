package environment

import (
	"os"
	"path/filepath"
)

var (
	WorkingDir, _ = filepath.Abs(".")
	ProgramDir    = filepath.Dir(os.Args[0])
)
