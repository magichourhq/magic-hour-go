package types

// The create endpoint to price, named by its path. For example `ai-image-editor` for `POST /v1/ai-image-editor`.
type V1EstimateImageListToolEnum string

const (
	V1EstimateImageListToolEnumAiClothesChanger       V1EstimateImageListToolEnum = "ai-clothes-changer"
	V1EstimateImageListToolEnumAiFaceEditor           V1EstimateImageListToolEnum = "ai-face-editor"
	V1EstimateImageListToolEnumAiGifGenerator         V1EstimateImageListToolEnum = "ai-gif-generator"
	V1EstimateImageListToolEnumAiHeadshotGenerator    V1EstimateImageListToolEnum = "ai-headshot-generator"
	V1EstimateImageListToolEnumAiImageEditor          V1EstimateImageListToolEnum = "ai-image-editor"
	V1EstimateImageListToolEnumAiImageGenerator       V1EstimateImageListToolEnum = "ai-image-generator"
	V1EstimateImageListToolEnumAiImageUpscaler        V1EstimateImageListToolEnum = "ai-image-upscaler"
	V1EstimateImageListToolEnumAiMemeGenerator        V1EstimateImageListToolEnum = "ai-meme-generator"
	V1EstimateImageListToolEnumAiQrCodeGenerator      V1EstimateImageListToolEnum = "ai-qr-code-generator"
	V1EstimateImageListToolEnumBodySwap               V1EstimateImageListToolEnum = "body-swap"
	V1EstimateImageListToolEnumFaceSwapPhoto          V1EstimateImageListToolEnum = "face-swap-photo"
	V1EstimateImageListToolEnumHeadSwap               V1EstimateImageListToolEnum = "head-swap"
	V1EstimateImageListToolEnumImageBackgroundRemover V1EstimateImageListToolEnum = "image-background-remover"
	V1EstimateImageListToolEnumPhotoColorizer         V1EstimateImageListToolEnum = "photo-colorizer"
)
