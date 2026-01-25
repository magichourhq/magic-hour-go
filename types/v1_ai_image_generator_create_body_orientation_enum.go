package types

// DEPRECATED: Use `aspect_ratio` instead.
//
// The orientation of the output image(s). `aspect_ratio` takes precedence when `orientation` if both are provided.
type V1AiImageGeneratorCreateBodyOrientationEnum string

const (
	V1AiImageGeneratorCreateBodyOrientationEnumLandscape V1AiImageGeneratorCreateBodyOrientationEnum = "landscape"
	V1AiImageGeneratorCreateBodyOrientationEnumPortrait  V1AiImageGeneratorCreateBodyOrientationEnum = "portrait"
	V1AiImageGeneratorCreateBodyOrientationEnumSquare    V1AiImageGeneratorCreateBodyOrientationEnum = "square"
)
