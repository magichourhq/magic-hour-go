package types

// * `default` - Use the default recommended prompt for the art style.
// * `custom` - Only use the prompt passed in the API. Note: for v1, lora prompt will still be auto added to apply the art style properly.
// * `append_default` - Add the default recommended prompt to the end of the prompt passed in the API.
type PostV1VideoToVideoBodyStylePromptTypeEnum string

const (
	PostV1VideoToVideoBodyStylePromptTypeEnumAppendDefault PostV1VideoToVideoBodyStylePromptTypeEnum = "append_default"
	PostV1VideoToVideoBodyStylePromptTypeEnumCustom        PostV1VideoToVideoBodyStylePromptTypeEnum = "custom"
	PostV1VideoToVideoBodyStylePromptTypeEnumDefault       PostV1VideoToVideoBodyStylePromptTypeEnum = "default"
)
