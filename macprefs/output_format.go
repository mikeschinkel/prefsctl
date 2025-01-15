package macprefs

type OutputFormat string

const (
	TXTFormat  OutputFormat = "txt"
	YAMLFormat OutputFormat = "yaml"
	JSONFormat OutputFormat = "json"
)

var AllFormats = []OutputFormat{
	TXTFormat,
	YAMLFormat,
	JSONFormat,
}
