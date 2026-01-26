package types

// Controls overall motion style.
// * `realistic` - Maintains likeness well, high quality, and reliable.
// * `prompted` - Slightly lower likeness; allows option to prompt scene.
//
// **Deprecated values (maintained for backward compatibility):**
// * `pro` - Deprecated: use `realistic`
// * `standard` - Deprecated: use `prompted`
// * `stable` - Deprecated: use `realistic`
// * `expressive` - Deprecated: use `prompted`
type V1AiTalkingPhotoCreateBodyStyleGenerationModeEnum string

const (
	V1AiTalkingPhotoCreateBodyStyleGenerationModeEnumExpressive V1AiTalkingPhotoCreateBodyStyleGenerationModeEnum = "expressive"
	V1AiTalkingPhotoCreateBodyStyleGenerationModeEnumPro        V1AiTalkingPhotoCreateBodyStyleGenerationModeEnum = "pro"
	V1AiTalkingPhotoCreateBodyStyleGenerationModeEnumPrompted   V1AiTalkingPhotoCreateBodyStyleGenerationModeEnum = "prompted"
	V1AiTalkingPhotoCreateBodyStyleGenerationModeEnumRealistic  V1AiTalkingPhotoCreateBodyStyleGenerationModeEnum = "realistic"
	V1AiTalkingPhotoCreateBodyStyleGenerationModeEnumStable     V1AiTalkingPhotoCreateBodyStyleGenerationModeEnum = "stable"
	V1AiTalkingPhotoCreateBodyStyleGenerationModeEnumStandard   V1AiTalkingPhotoCreateBodyStyleGenerationModeEnum = "standard"
)
