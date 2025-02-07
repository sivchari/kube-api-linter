package config

// LintersConfig contains configuration for individual linters.
type LintersConfig struct {
	// conditions contains configuration for the conditions linter.
	Conditions ConditionsConfig `json:"conditions"`

	// jsonTags contains configuration for the jsontags linter.
	JSONTags JSONTagsConfig `json:"jsonTags"`

	// optionalOrRequired contains configuration for the optionalorrequired linter.
	OptionalOrRequired OptionalOrRequiredConfig `json:"optionalOrRequired"`

	// requiredFields contains configuration for the requiredfields linter.
	RequiredFields RequiredFieldsConfig `json:"requiredFields"`
}

// ConditionsFirstField is the policy for the conditions linter.
type ConditionsFirstField string

const (
	// ConditionsFirstFieldWarn indicates that the conditions should be the first field in the struct.
	ConditionsFirstFieldWarn ConditionsFirstField = "Warn"

	// ConditionsFirstFieldIgnore indicates that the conditions do not need to be the first field in the struct.
	ConditionsFirstFieldIgnore ConditionsFirstField = "Ignore"
)

// ConditionsUseProtobuf is the policy for the conditions linter.
type ConditionsUseProtobuf string

const (
	// ConditionsUseProtobufSuggestFix indicates that the linter will emit a warning if the conditions are not using protobuf tags and suggest a fix.
	ConditionsUseProtobufSuggestFix ConditionsUseProtobuf = "SuggestFix"

	// ConditionsUseProtobufWarn indicates that the linter will emit a warning if the conditions are not using protobuf tags.
	ConditionsUseProtobufWarn ConditionsUseProtobuf = "Warn"

	// ConditionsUseProtobufIgnore indicates that the linter will not emit a warning if the conditions are not using protobuf tags.
	ConditionsUseProtobufIgnore ConditionsUseProtobuf = "Ignore"

	// ConditionsUseProtobufForbid indicates that the linter will emit an error if the conditions are using protobuf tags, a fix will also be suggested.
	ConditionsUseProtobufForbid ConditionsUseProtobuf = "Forbid"
)

// ConditionsUsePatchStrategy is the policy for the conditions linter.
type ConditionsUsePatchStrategy string

const (
	// ConditionsUsePatchStrategySuggestFix indicates that the linter will emit a warning if the conditions are not using patch strategy tags and suggest a fix.
	ConditionsUsePatchStrategySuggestFix ConditionsUsePatchStrategy = "SuggestFix"

	// ConditionsUsePatchStrategyWarn indicates that the linter will emit a warning if the conditions are not using patch strategy tags.
	ConditionsUsePatchStrategyWarn ConditionsUsePatchStrategy = "Warn"

	// ConditionsUsePatchStrategyIgnore indicates that the linter will not emit a warning if the conditions are not using patch strategy tags.
	ConditionsUsePatchStrategyIgnore ConditionsUsePatchStrategy = "Ignore"

	// ConditionsUsePatchStrategyForbid indicates that the linter will emit an error if the conditions are using patch strategy tags, a fix will also be suggested.
	ConditionsUsePatchStrategyForbid ConditionsUsePatchStrategy = "Forbid"
)

// ConditionsConfig contains configuration for the conditions linter.
type ConditionsConfig struct {
	// isFirstField indicates whether the conditions should be the first field in the struct.
	// Valid values are Warn and Ignore.
	// When set to Warn, the linter will emit a warning if the conditions are not the first field in the struct.
	// When set to Ignore, the linter will not emit a warning if the conditions are not the first field in the struct.
	// When otherwise not specified, the default value is Warn.
	IsFirstField ConditionsFirstField `json:"isFirstField"`

	// useProtobuf indicates whether the linter should use protobuf tags.
	// Valid values are SuggestFix, Warn, Ignore and Forbid.
	// When set to SuggestFix, the linter will emit a warning if the conditions are not using protobuf tags and suggest a fix.
	// When set to Warn, the linter will emit a warning if the conditions are not using protobuf tags.
	// When set to Ignore, the linter will not emit a warning if the conditions are not using protobuf tags.
	// When set to Forbid, the linter will emit an error if the conditions are using protobuf tags, a fix will also be suggested.
	// When otherwise not specified, the default value is SuggestFix.
	UseProtobuf ConditionsUseProtobuf `json:"useProtobuf"`

	// usePatchStrategy indicates whether the linter should enforce the patch strategy tags.
	// Valid values are SuggestFix, Warn, Ignore and Forbid.
	// When set to SuggestFix, the linter will emit a warning if the conditions are not using patch strategy tags and suggest a fix.
	// When set to Warn, the linter will emit a warning if the conditions are not using patch strategy tags.
	// When set to Ignore, the linter will not emit a warning if the conditions are not using patch strategy tags.
	// When set to Forbid, the linter will emit an error if the conditions are using patch strategy tags, a fix will also be suggested.
	// When otherwise not specified, the default value is SuggestFix.
	UsePatchStrategy ConditionsUsePatchStrategy `json:"usePatchStrategy"`
}

// JSONTagsConfig contains configuration for the jsontags linter.
type JSONTagsConfig struct {
	// jsonTagRegex is the regular expression used to validate that json tags are in a particular format.
	// By default, the regex used is "^[a-z][a-z0-9]*(?:[A-Z][a-z0-9]*)*$" and is used to check for
	// camel case like string.
	JSONTagRegex string `json:"jsonTagRegex"`
}

// OptionalOrRequiredConfig contains configuration for the optionalorrequired linter.
type OptionalOrRequiredConfig struct {
	// preferredOptionalMarker is the preferred marker to use for optional fields.
	// If this field is not set, the default value is "optional".
	// Valid values are "optional" and "kubebuilder:validation:Optional".
	PreferredOptionalMarker string `json:"preferredOptionalMarker"`

	// preferredRequiredMarker is the preferred marker to use for required fields.
	// If this field is not set, the default value is "required".
	// Valid values are "required" and "kubebuilder:validation:Required".
	PreferredRequiredMarker string `json:"preferredRequiredMarker"`
}

// RequiredFieldPointerPolicy is the policy for pointers in required fields.
type RequiredFieldPointerPolicy string

const (
	// RequiredFieldPointerWarn indicates that the linter will emit a warning if a required field is a pointer.
	RequiredFieldPointerWarn RequiredFieldPointerPolicy = "Warn"

	// RequiredFieldPointerSuggestFix indicates that the linter will emit a warning if a required field is a pointer and suggest a fix.
	RequiredFieldPointerSuggestFix RequiredFieldPointerPolicy = "SuggestFix"
)

// RequiredFieldsConfig contains configuration for the requiredfields linter.
type RequiredFieldsConfig struct {
	// pointerPolicy is the policy for pointers in required fields.
	// Valid values are "Warn" and "SuggestFix".
	// When set to "Warn", the linter will emit a warning if a required field is a pointer.
	// When set to "SuggestFix", the linter will emit a warning if a required field is a pointer and suggest a fix.
	// When otherwise not specified, the default value is "SuggestFix".
	PointerPolicy RequiredFieldPointerPolicy `json:"pointerPolicy"`
}
