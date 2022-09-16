package embedfiles

import _ "embed"

var (
	//go:embed files/edoex_logo.txt
	EdoexLogo string

	//go:embed files/default_expansion_config.yaml
	DefaultExpansionConfig string

	//go:embed files/create_tables.sql
	CreateTablesScript string

	//go:embed files/default_script.lua
	DefaultCardScript string

	//go:embed files/cards_readme.md
	CardsReadme string

	//go:embed files/meta_readme.md
	MetaReadme string

	//go:embed files/scripts_readme.md
	ScriptsReadme string

	//go:embed files/artworks_readme.md
	ArtworksReadme string
)
