package applemdm

// YAMLSchema represents the top-level schema object
//
//goland:noinspection SpellCheckingInspection
type YAMLSchema struct {
	Title        string           `yaml:"title"`
	Description  string           `yaml:"description"`
	Payload      *YAMLPayload     `yaml:"payload"`
	PayloadKeys  []YAMLPayloadKey `yaml:"payloadkeys,omitempty"`
	ResponseKeys []YAMLPayloadKey `yaml:"responsekeys,omitempty"`
	Reasons      []YAMLReasonCode `yaml:"reasons,omitempty"`
	Notes        []string         `yaml:"notes,omitempty"`
}

// YAMLPayload represents information about the object as a whole
//
//goland:noinspection SpellCheckingInspection
type YAMLPayload struct {
	PayloadType     string         `yaml:"payloadtype,omitempty"`
	RequestType     string         `yaml:"requesttype,omitempty"`
	DeclarationType string         `yaml:"declarationtype,omitempty"`
	StatusItemType  string         `yaml:"statusitemtype,omitempty"`
	CredentialType  string         `yaml:"credentialtype,omitempty"`
	SupportedOS     *YAMLOSSupport `yaml:"supportedOS,omitempty"`
	Apply           string         `yaml:"apply,omitempty"`
	Content         string         `yaml:"content,omitempty"`
}

// YAMLOSSupport defines supported features per operating system
type YAMLOSSupport struct {
	IOS      *YAMLOSFeatures `yaml:"iOS,omitempty"`
	MacOS    *YAMLOSFeatures `yaml:"macOS,omitempty"`
	TVOS     *YAMLOSFeatures `yaml:"tvOS,omitempty"`
	VisionOS *YAMLOSFeatures `yaml:"visionOS,omitempty"`
	WatchOS  *YAMLOSFeatures `yaml:"watchOS,omitempty"`
}

// YAMLOSFeatures defines the supported features and requirements for an OS
//
//goland:noinspection SpellCheckingInspection
type YAMLOSFeatures struct {
	Introduced         string              `yaml:"introduced,omitempty"`
	Deprecated         string              `yaml:"deprecated,omitempty"`
	Removed            string              `yaml:"removed,omitempty"`
	AccessRights       string              `yaml:"accessrights,omitempty"`
	Multiple           bool                `yaml:"multiple,omitempty"`
	DeviceChannel      bool                `yaml:"devicechannel,omitempty"`
	UserChannel        bool                `yaml:"userchannel,omitempty"`
	Supervised         bool                `yaml:"supervised,omitempty"`
	RequiresDEP        bool                `yaml:"requiresdep,omitempty"`
	UserApprovedMDM    bool                `yaml:"userapprovedmdm,omitempty"`
	AllowManualInstall bool                `yaml:"allowmanualinstall,omitempty"`
	SharedIPad         *YAMLSharedIPad     `yaml:"sharedipad,omitempty"`
	UserEnrollment     *YAMLUserEnrollment `yaml:"userenrollment,omitempty"`
	AlwaysSkippable    bool                `yaml:"always-skippable,omitempty"`
	AllowedEnrollments []string            `yaml:"allowed-enrollments,omitempty"`
	AllowedScopes      []string            `yaml:"allowed-scopes,omitempty"`
}

// YAMLSharedIPad defines shared iPad specific features
//
//goland:noinspection SpellCheckingInspection
type YAMLSharedIPad struct {
	Mode          string   `yaml:"mode"`
	DeviceChannel bool     `yaml:"devicechannel,omitempty"`
	UserChannel   bool     `yaml:"userchannel,omitempty"`
	AllowedScopes []string `yaml:"allowed-scopes,omitempty"`
}

// YAMLUserEnrollment defines user enrollment specific features
type YAMLUserEnrollment struct {
	Mode     string `yaml:"mode"`
	Behavior string `yaml:"behavior,omitempty"`
}

// YAMLPayloadKey represents a key in the payload or response
//
//goland:noinspection SpellCheckingInspection
type YAMLPayloadKey struct {
	Key         string           `yaml:"key"`
	Title       string           `yaml:"title,omitempty"`
	SupportedOS *YAMLOSSupport   `yaml:"supportedOS,omitempty"`
	Type        string           `yaml:"type"`
	Subtype     string           `yaml:"subtype,omitempty"`
	ValueType   string           `yaml:"valuetype,omitempty"`
	AssetTypes  []string         `yaml:"assettypes,omitempty"`
	Presence    string           `yaml:"presence,omitempty"`
	RangeList   []string         `yaml:"rangelist,omitempty"`
	Range       *YAMLRange       `yaml:"range,omitempty"`
	Default     interface{}      `yaml:"default,omitempty"`
	Format      string           `yaml:"format,omitempty"`
	Repetition  *YAMLRepetition  `yaml:"repetition,omitempty"`
	CombineType string           `yaml:"combinetype,omitempty"`
	Content     string           `yaml:"content,omitempty"`
	SubKeyType  string           `yaml:"subkeytype,omitempty"`
	SubKeys     []YAMLPayloadKey `yaml:"subkeys,omitempty"`
}

// YAMLRange defines numeric bounds
type YAMLRange struct {
	Min float64 `yaml:"min"`
	Max float64 `yaml:"max"`
}

// YAMLRepetition defines cardinality bounds
type YAMLRepetition struct {
	Min float64 `yaml:"min"`
	Max float64 `yaml:"max"`
}

// YAMLReasonCode represents a declarative device management status reason code
type YAMLReasonCode struct {
	Code        string `yaml:"code"`
	Description string `yaml:"description"`
}
