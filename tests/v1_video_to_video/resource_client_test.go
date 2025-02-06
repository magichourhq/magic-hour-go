package test_video_to_video_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	video_to_video "github.com/magichourhq/magic-hour-go/resources/v1/video_to_video"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithBearerAuth("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/magichour/magic-hour/0.8.4"))
	res, err := client.V1.VideoToVideo.Create(video_to_video.CreateRequest{Assets: types.PostV1VideoToVideoBodyAssets{VideoFilePath: nullable.NewValue("video/id/1234.mp4"), VideoSource: types.PostV1VideoToVideoBodyAssetsVideoSourceEnumFile}, EndSeconds: 15.0, Height: 960, StartSeconds: 0.0, Style: types.PostV1VideoToVideoBodyStyle{ArtStyle: types.PostV1VideoToVideoBodyStyleArtStyleEnum3dRender, Model: types.PostV1VideoToVideoBodyStyleModelEnumAbsoluteReality, Prompt: nullable.NewNull[string](), PromptType: types.PostV1VideoToVideoBodyStylePromptTypeEnumAppendDefault, Version: types.PostV1VideoToVideoBodyStyleVersionEnumDefault}, Width: 512})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
