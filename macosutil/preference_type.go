package macosutil

type PreferenceType string

var (
	InvalidType PreferenceType = "invalid"
	StringType  PreferenceType = "string"
	IntegerType PreferenceType = "integer"
	BooleanType PreferenceType = "boolean"
	ArrayType   PreferenceType = "array"
)

//var (
//	UnknownType  PreferenceType = "unknownType"
//	StringType   PreferenceType = "stringType"
//	NumberType   PreferenceType = "numberType"
//	IntType      PreferenceType = "intType"
//	FloatType    PreferenceType = "floatType"
//	BooleanType     PreferenceType = "boolType"
//	IntBoolType  PreferenceType = "intBoolType"
//	LanguageType PreferenceType = "languageType"
//	LocaleType   PreferenceType = "localeType"
//)
