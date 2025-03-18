package v1

import (
	http "net/http"

	sdkcore "github.com/magichourhq/magic-hour-go/core"
	ai_clothes_changer "github.com/magichourhq/magic-hour-go/resources/v1/ai_clothes_changer"
	ai_headshot_generator "github.com/magichourhq/magic-hour-go/resources/v1/ai_headshot_generator"
	ai_image_generator "github.com/magichourhq/magic-hour-go/resources/v1/ai_image_generator"
	ai_image_upscaler "github.com/magichourhq/magic-hour-go/resources/v1/ai_image_upscaler"
	ai_photo_editor "github.com/magichourhq/magic-hour-go/resources/v1/ai_photo_editor"
	ai_qr_code_generator "github.com/magichourhq/magic-hour-go/resources/v1/ai_qr_code_generator"
	animation "github.com/magichourhq/magic-hour-go/resources/v1/animation"
	face_swap "github.com/magichourhq/magic-hour-go/resources/v1/face_swap"
	face_swap_photo "github.com/magichourhq/magic-hour-go/resources/v1/face_swap_photo"
	files "github.com/magichourhq/magic-hour-go/resources/v1/files"
	image_background_remover "github.com/magichourhq/magic-hour-go/resources/v1/image_background_remover"
	image_projects "github.com/magichourhq/magic-hour-go/resources/v1/image_projects"
	image_to_video "github.com/magichourhq/magic-hour-go/resources/v1/image_to_video"
	lip_sync "github.com/magichourhq/magic-hour-go/resources/v1/lip_sync"
	text_to_video "github.com/magichourhq/magic-hour-go/resources/v1/text_to_video"
	video_projects "github.com/magichourhq/magic-hour-go/resources/v1/video_projects"
	video_to_video "github.com/magichourhq/magic-hour-go/resources/v1/video_to_video"
)

type Client struct {
	coreClient             *sdkcore.CoreClient
	ImageProjects          *image_projects.Client
	VideoProjects          *video_projects.Client
	AiClothesChanger       *ai_clothes_changer.Client
	AiHeadshotGenerator    *ai_headshot_generator.Client
	AiImageGenerator       *ai_image_generator.Client
	AiImageUpscaler        *ai_image_upscaler.Client
	AiPhotoEditor          *ai_photo_editor.Client
	AiQrCodeGenerator      *ai_qr_code_generator.Client
	Animation              *animation.Client
	FaceSwap               *face_swap.Client
	FaceSwapPhoto          *face_swap_photo.Client
	Files                  *files.Client
	ImageBackgroundRemover *image_background_remover.Client
	ImageToVideo           *image_to_video.Client
	LipSync                *lip_sync.Client
	TextToVideo            *text_to_video.Client
	VideoToVideo           *video_to_video.Client
}
type RequestModifier = func(req *http.Request) error

// Instantiate a new resource client
func NewClient(coreClient *sdkcore.CoreClient) *Client {
	client := Client{
		coreClient:             coreClient,
		ImageProjects:          image_projects.NewClient(coreClient),
		VideoProjects:          video_projects.NewClient(coreClient),
		AiClothesChanger:       ai_clothes_changer.NewClient(coreClient),
		AiHeadshotGenerator:    ai_headshot_generator.NewClient(coreClient),
		AiImageGenerator:       ai_image_generator.NewClient(coreClient),
		AiImageUpscaler:        ai_image_upscaler.NewClient(coreClient),
		AiPhotoEditor:          ai_photo_editor.NewClient(coreClient),
		AiQrCodeGenerator:      ai_qr_code_generator.NewClient(coreClient),
		Animation:              animation.NewClient(coreClient),
		FaceSwap:               face_swap.NewClient(coreClient),
		FaceSwapPhoto:          face_swap_photo.NewClient(coreClient),
		Files:                  files.NewClient(coreClient),
		ImageBackgroundRemover: image_background_remover.NewClient(coreClient),
		ImageToVideo:           image_to_video.NewClient(coreClient),
		LipSync:                lip_sync.NewClient(coreClient),
		TextToVideo:            text_to_video.NewClient(coreClient),
		VideoToVideo:           video_to_video.NewClient(coreClient),
	}

	return &client
}
