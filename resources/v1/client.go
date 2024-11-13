// Code generated by Sideko (sideko.dev) DO NOT EDIT.

package v1

import (
	http "net/http"

	sdkcore "github.com/magichourhq/magic-hour-go/core"
	face_swap "github.com/magichourhq/magic-hour-go/resources/v1/face_swap"
	face_swap_photo "github.com/magichourhq/magic-hour-go/resources/v1/face_swap_photo"
	files "github.com/magichourhq/magic-hour-go/resources/v1/files"
	image_projects "github.com/magichourhq/magic-hour-go/resources/v1/image_projects"
	image_to_video "github.com/magichourhq/magic-hour-go/resources/v1/image_to_video"
	lip_sync "github.com/magichourhq/magic-hour-go/resources/v1/lip_sync"
	text_to_video "github.com/magichourhq/magic-hour-go/resources/v1/text_to_video"
	video_projects "github.com/magichourhq/magic-hour-go/resources/v1/video_projects"
	video_to_video "github.com/magichourhq/magic-hour-go/resources/v1/video_to_video"
)

type Client struct {
	coreClient    *sdkcore.CoreClient
	ImageProjects *image_projects.Client
	VideoProjects *video_projects.Client
	FaceSwap      *face_swap.Client
	FaceSwapPhoto *face_swap_photo.Client
	Files         *files.Client
	ImageToVideo  *image_to_video.Client
	LipSync       *lip_sync.Client
	TextToVideo   *text_to_video.Client
	VideoToVideo  *video_to_video.Client
}
type RequestModifier = func(req *http.Request) error

// Instantiate a new resource client
func NewClient(coreClient *sdkcore.CoreClient) *Client {
	client := Client{
		coreClient:    coreClient,
		ImageProjects: image_projects.NewClient(coreClient),
		VideoProjects: video_projects.NewClient(coreClient),
		FaceSwap:      face_swap.NewClient(coreClient),
		FaceSwapPhoto: face_swap_photo.NewClient(coreClient),
		Files:         files.NewClient(coreClient),
		ImageToVideo:  image_to_video.NewClient(coreClient),
		LipSync:       lip_sync.NewClient(coreClient),
		TextToVideo:   text_to_video.NewClient(coreClient),
		VideoToVideo:  video_to_video.NewClient(coreClient),
	}

	return &client
}