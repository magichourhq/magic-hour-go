package v1

import (
	http "net/http"

	sdkcore "github.com/magichourhq/magic-hour-go/core"
	ai_clothes_changer "github.com/magichourhq/magic-hour-go/resources/v1/ai_clothes_changer"
	ai_face_editor "github.com/magichourhq/magic-hour-go/resources/v1/ai_face_editor"
	ai_gif_generator "github.com/magichourhq/magic-hour-go/resources/v1/ai_gif_generator"
	ai_headshot_generator "github.com/magichourhq/magic-hour-go/resources/v1/ai_headshot_generator"
	ai_image_editor "github.com/magichourhq/magic-hour-go/resources/v1/ai_image_editor"
	ai_image_generator "github.com/magichourhq/magic-hour-go/resources/v1/ai_image_generator"
	ai_image_upscaler "github.com/magichourhq/magic-hour-go/resources/v1/ai_image_upscaler"
	ai_meme_generator "github.com/magichourhq/magic-hour-go/resources/v1/ai_meme_generator"
	ai_photo_editor "github.com/magichourhq/magic-hour-go/resources/v1/ai_photo_editor"
	ai_qr_code_generator "github.com/magichourhq/magic-hour-go/resources/v1/ai_qr_code_generator"
	ai_talking_photo "github.com/magichourhq/magic-hour-go/resources/v1/ai_talking_photo"
	ai_voice_generator "github.com/magichourhq/magic-hour-go/resources/v1/ai_voice_generator"
	animation "github.com/magichourhq/magic-hour-go/resources/v1/animation"
	audio_projects "github.com/magichourhq/magic-hour-go/resources/v1/audio_projects"
	auto_subtitle_generator "github.com/magichourhq/magic-hour-go/resources/v1/auto_subtitle_generator"
	face_detection "github.com/magichourhq/magic-hour-go/resources/v1/face_detection"
	face_swap "github.com/magichourhq/magic-hour-go/resources/v1/face_swap"
	face_swap_photo "github.com/magichourhq/magic-hour-go/resources/v1/face_swap_photo"
	files "github.com/magichourhq/magic-hour-go/resources/v1/files"
	image_background_remover "github.com/magichourhq/magic-hour-go/resources/v1/image_background_remover"
	image_projects "github.com/magichourhq/magic-hour-go/resources/v1/image_projects"
	image_to_video "github.com/magichourhq/magic-hour-go/resources/v1/image_to_video"
	lip_sync "github.com/magichourhq/magic-hour-go/resources/v1/lip_sync"
	photo_colorizer "github.com/magichourhq/magic-hour-go/resources/v1/photo_colorizer"
	text_to_video "github.com/magichourhq/magic-hour-go/resources/v1/text_to_video"
	video_projects "github.com/magichourhq/magic-hour-go/resources/v1/video_projects"
	video_to_video "github.com/magichourhq/magic-hour-go/resources/v1/video_to_video"
)

type Client struct {
	coreClient             *sdkcore.CoreClient
	AudioProjects          *audio_projects.Client
	ImageProjects          *image_projects.Client
	VideoProjects          *video_projects.Client
	FaceDetection          *face_detection.Client
	AiClothesChanger       *ai_clothes_changer.Client
	AiFaceEditor           *ai_face_editor.Client
	AiGifGenerator         *ai_gif_generator.Client
	AiHeadshotGenerator    *ai_headshot_generator.Client
	AiImageEditor          *ai_image_editor.Client
	AiImageGenerator       *ai_image_generator.Client
	AiImageUpscaler        *ai_image_upscaler.Client
	AiMemeGenerator        *ai_meme_generator.Client
	AiPhotoEditor          *ai_photo_editor.Client
	AiQrCodeGenerator      *ai_qr_code_generator.Client
	AiTalkingPhoto         *ai_talking_photo.Client
	AiVoiceGenerator       *ai_voice_generator.Client
	Animation              *animation.Client
	AutoSubtitleGenerator  *auto_subtitle_generator.Client
	FaceSwap               *face_swap.Client
	FaceSwapPhoto          *face_swap_photo.Client
	Files                  *files.Client
	ImageBackgroundRemover *image_background_remover.Client
	ImageToVideo           *image_to_video.Client
	LipSync                *lip_sync.Client
	PhotoColorizer         *photo_colorizer.Client
	TextToVideo            *text_to_video.Client
	VideoToVideo           *video_to_video.Client
}
type RequestModifier = func(req *http.Request) error

// Instantiate a new resource client
func NewClient(coreClient *sdkcore.CoreClient) *Client {
	client := Client{
		coreClient:             coreClient,
		AudioProjects:          audio_projects.NewClient(coreClient),
		ImageProjects:          image_projects.NewClient(coreClient),
		VideoProjects:          video_projects.NewClient(coreClient),
		FaceDetection:          face_detection.NewClient(coreClient),
		AiClothesChanger:       ai_clothes_changer.NewClient(coreClient),
		AiFaceEditor:           ai_face_editor.NewClient(coreClient),
		AiGifGenerator:         ai_gif_generator.NewClient(coreClient),
		AiHeadshotGenerator:    ai_headshot_generator.NewClient(coreClient),
		AiImageEditor:          ai_image_editor.NewClient(coreClient),
		AiImageGenerator:       ai_image_generator.NewClient(coreClient),
		AiImageUpscaler:        ai_image_upscaler.NewClient(coreClient),
		AiMemeGenerator:        ai_meme_generator.NewClient(coreClient),
		AiPhotoEditor:          ai_photo_editor.NewClient(coreClient),
		AiQrCodeGenerator:      ai_qr_code_generator.NewClient(coreClient),
		AiTalkingPhoto:         ai_talking_photo.NewClient(coreClient),
		AiVoiceGenerator:       ai_voice_generator.NewClient(coreClient),
		Animation:              animation.NewClient(coreClient),
		AutoSubtitleGenerator:  auto_subtitle_generator.NewClient(coreClient),
		FaceSwap:               face_swap.NewClient(coreClient),
		FaceSwapPhoto:          face_swap_photo.NewClient(coreClient),
		Files:                  files.NewClient(coreClient),
		ImageBackgroundRemover: image_background_remover.NewClient(coreClient),
		ImageToVideo:           image_to_video.NewClient(coreClient),
		LipSync:                lip_sync.NewClient(coreClient),
		PhotoColorizer:         photo_colorizer.NewClient(coreClient),
		TextToVideo:            text_to_video.NewClient(coreClient),
		VideoToVideo:           video_to_video.NewClient(coreClient),
	}

	return &client
}
