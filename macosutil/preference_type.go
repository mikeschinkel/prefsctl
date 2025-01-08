package macosutil

type PreferenceType string

var (
	UnknownType  PreferenceType = "unknownType"
	StringType   PreferenceType = "stringType"
	NumberType   PreferenceType = "numberType"
	IntType      PreferenceType = "intType"
	FloatType    PreferenceType = "floatType"
	BoolType     PreferenceType = "boolType"
	IntBoolType  PreferenceType = "intBoolType"
	LanguageType PreferenceType = "languageType"
	LocaleType   PreferenceType = "localeType"
)
