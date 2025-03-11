package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	video_to_video "github.com/magichourhq/magic-hour-go/resources/v1/video_to_video"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.VideoToVideo.Create(video_to_video.CreateRequest{
		Assets: types.PostV1VideoToVideoBodyAssets{
			VideoFilePath: nullable.NewValue("api-assets/id/1234.mp4"),
			VideoSource:   types.PostV1VideoToVideoBodyAssetsVideoSourceEnumFile,
		},
		EndSeconds:   15.0,
		Height:       960,
		StartSeconds: 0.0,
		Style: types.PostV1VideoToVideoBodyStyle{
			ArtStyle:   types.PostV1VideoToVideoBodyStyleArtStyleEnum3dRender,
			Model:      types.PostV1VideoToVideoBodyStyleModelEnumAbsoluteReality,
			Prompt:     nullable.NewNull[string](),
			PromptType: types.PostV1VideoToVideoBodyStylePromptTypeEnumAppendDefault,
			Version:    types.PostV1VideoToVideoBodyStyleVersionEnumDefault,
		},
		Width: 512,
	})
}
