package environment

import "os"

func ClearBuild() error {
	for _, p := range []string{
		BuildStringsPath(), BuildDatabasePath(), BuildPicsPath(),
		BuildScriptsPath(),
	} {
		err := os.RemoveAll(p)

		if err != nil {
			return err
		}
	}

	return nil
}
