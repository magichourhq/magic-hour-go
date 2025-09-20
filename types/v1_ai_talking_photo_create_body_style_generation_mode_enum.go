package types

// Controls overall motion style.
// * `pro` -  Higher fidelity, realistic detail, accurate lip sync, and faster generation.
// * `standard` -  More expressive motion, but lower visual fidelity.
//
// * `expressive` - More motion and facial expressiveness; may introduce visual artifacts. (Deprecated: passing this value will be treated as `standard`)
// * `stable` -  Reduced motion for cleaner output; may result in minimal animation. (Deprecated: passing this value will be treated as `pro`)
type V1AiTalkingPhotoCreateBodyStyleGenerationModeEnum string

const (
	V1AiTalkingPhotoCreateBodyStyleGenerationModeEnumExpressive V1AiTalkingPhotoCreateBodyStyleGenerationModeEnum = "expressive"
	V1AiTalkingPhotoCreateBodyStyleGenerationModeEnumPro        V1AiTalkingPhotoCreateBodyStyleGenerationModeEnum = "pro"
	V1AiTalkingPhotoCreateBodyStyleGenerationModeEnumStable     V1AiTalkingPhotoCreateBodyStyleGenerationModeEnum = "stable"
	V1AiTalkingPhotoCreateBodyStyleGenerationModeEnumStandard   V1AiTalkingPhotoCreateBodyStyleGenerationModeEnum = "standard"
)
