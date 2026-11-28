package types

// The create endpoint to price, named by its path. For example `image-to-video` for `POST /v1/image-to-video`.
type V1EstimateVideoListToolEnum string

const (
	V1EstimateVideoListToolEnumAiUgcAd               V1EstimateVideoListToolEnum = "ai-ugc-ad"
	V1EstimateVideoListToolEnumAiVideoEditor         V1EstimateVideoListToolEnum = "ai-video-editor"
	V1EstimateVideoListToolEnumAiVideoTranslator     V1EstimateVideoListToolEnum = "ai-video-translator"
	V1EstimateVideoListToolEnumAiVideoUpscaler       V1EstimateVideoListToolEnum = "ai-video-upscaler"
	V1EstimateVideoListToolEnumAudioToVideo          V1EstimateVideoListToolEnum = "audio-to-video"
	V1EstimateVideoListToolEnumAutoSubtitleGenerator V1EstimateVideoListToolEnum = "auto-subtitle-generator"
	V1EstimateVideoListToolEnumCharacterReplace      V1EstimateVideoListToolEnum = "character-replace"
	V1EstimateVideoListToolEnumFaceSwap              V1EstimateVideoListToolEnum = "face-swap"
	V1EstimateVideoListToolEnumImageToVideo          V1EstimateVideoListToolEnum = "image-to-video"
	V1EstimateVideoListToolEnumTextToVideo           V1EstimateVideoListToolEnum = "text-to-video"
)
