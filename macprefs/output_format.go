package macprefs

type OutputFormat string

const (
	TxtFormat  OutputFormat = "txt"
	YAMLFormat OutputFormat = "yaml"
	JSONFormat OutputFormat = "json"
	GoFormat   OutputFormat = "go"
)
