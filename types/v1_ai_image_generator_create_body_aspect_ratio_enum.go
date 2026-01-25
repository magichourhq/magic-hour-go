package types

// The aspect ratio of the output image(s). If not specified, defaults to `1:1` (square).
type V1AiImageGeneratorCreateBodyAspectRatioEnum string

const (
	V1AiImageGeneratorCreateBodyAspectRatioEnum169 V1AiImageGeneratorCreateBodyAspectRatioEnum = "16:9"
	V1AiImageGeneratorCreateBodyAspectRatioEnum11  V1AiImageGeneratorCreateBodyAspectRatioEnum = "1:1"
	V1AiImageGeneratorCreateBodyAspectRatioEnum916 V1AiImageGeneratorCreateBodyAspectRatioEnum = "9:16"
)
