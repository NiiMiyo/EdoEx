package embedfiles

import _ "embed"

var (
	//go:embed files/edoex_logo.txt
	EdoexLogo string

	//go:embed files/default_expansion_config.yaml
	DefaultExpansionConfig string
)
