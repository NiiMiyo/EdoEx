package environment

import "os"

func ClearBuild() error {
	for _, p := range []string{
		EdoproStringsBuildPath(), EdoproDatabaseBuildPath(), EdoproImagesBuildPath(),
		EdoproScriptsBuildPath(), OmegaDatabaseBuildPath(), OmegaImagesBuildPath(),
		OmegaScriptsBuildPath(),
	} {
		err := os.RemoveAll(p)

		if err != nil {
			return err
		}
	}

	return nil
}
