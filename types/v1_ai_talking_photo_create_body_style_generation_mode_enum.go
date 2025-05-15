package types

// Controls overall motion style.
// * `expressive` - More motion and facial expressiveness; may introduce visual artifacts.
// * `stable` -  Reduced motion for cleaner output; may result in minimal animation.
type V1AiTalkingPhotoCreateBodyStyleGenerationModeEnum string

const (
	V1AiTalkingPhotoCreateBodyStyleGenerationModeEnumExpressive V1AiTalkingPhotoCreateBodyStyleGenerationModeEnum = "expressive"
	V1AiTalkingPhotoCreateBodyStyleGenerationModeEnumStable     V1AiTalkingPhotoCreateBodyStyleGenerationModeEnum = "stable"
)
