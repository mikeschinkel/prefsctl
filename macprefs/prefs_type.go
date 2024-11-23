package macprefs

type PrefType string

const (
	StringType   PrefType = "string"
	NumberType   PrefType = "number"
	BoolType     PrefType = "bool"
	LanguageType PrefType = "language"
	LocaleType   PrefType = "locale"
)
