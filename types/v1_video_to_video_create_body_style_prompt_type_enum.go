package types

// * `default` - Use the default recommended prompt for the art style.
// * `custom` - Only use the prompt passed in the API. Note: for v1, lora prompt will still be auto added to apply the art style properly.
// * `append_default` - Add the default recommended prompt to the end of the prompt passed in the API.
type V1VideoToVideoCreateBodyStylePromptTypeEnum string

const (
	V1VideoToVideoCreateBodyStylePromptTypeEnumAppendDefault V1VideoToVideoCreateBodyStylePromptTypeEnum = "append_default"
	V1VideoToVideoCreateBodyStylePromptTypeEnumCustom        V1VideoToVideoCreateBodyStylePromptTypeEnum = "custom"
	V1VideoToVideoCreateBodyStylePromptTypeEnumDefault       V1VideoToVideoCreateBodyStylePromptTypeEnum = "default"
)
