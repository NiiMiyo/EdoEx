package embedfiles

import _ "embed"

var (
	//go:embed files/edoex_logo.txt
	EdoexLogo string
)
