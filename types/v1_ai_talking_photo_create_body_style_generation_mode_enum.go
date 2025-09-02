package types

// Controls overall motion style.
// * `pro` -  Realistic, high fidelity, accurate lip sync, slower.
// * `expressive` - More motion and facial expressiveness; may introduce visual artifacts.
// * `stable` -  Reduced motion for cleaner output; may result in minimal animation. (Deprecated: passing this value will be treated as `pro`)
type V1AiTalkingPhotoCreateBodyStyleGenerationModeEnum string

const (
	V1AiTalkingPhotoCreateBodyStyleGenerationModeEnumExpressive V1AiTalkingPhotoCreateBodyStyleGenerationModeEnum = "expressive"
	V1AiTalkingPhotoCreateBodyStyleGenerationModeEnumPro        V1AiTalkingPhotoCreateBodyStyleGenerationModeEnum = "pro"
	V1AiTalkingPhotoCreateBodyStyleGenerationModeEnumStable     V1AiTalkingPhotoCreateBodyStyleGenerationModeEnum = "stable"
)
