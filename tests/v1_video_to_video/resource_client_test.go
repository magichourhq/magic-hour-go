package test_video_to_video_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	video_to_video "github.com/magichourhq/magic-hour-go/resources/v1/video_to_video"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestCreate200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.VideoToVideo.Create(video_to_video.CreateRequest{
		Assets: types.V1VideoToVideoCreateBodyAssets{
			VideoFilePath: nullable.NewValue("api-assets/id/1234.mp4"),
			VideoSource:   types.V1VideoToVideoCreateBodyAssetsVideoSourceEnumFile,
			YoutubeUrl:    nullable.NewValue("http://www.example.com"),
		},
		EndSeconds:    15.0,
		FpsResolution: nullable.NewValue(types.V1VideoToVideoCreateBodyFpsResolutionEnumHalf),
		Height:        nullable.NewValue(960),
		Name:          nullable.NewValue("Video To Video video"),
		StartSeconds:  0.0,
		Style: types.V1VideoToVideoCreateBodyStyle{
			ArtStyle:   types.V1VideoToVideoCreateBodyStyleArtStyleEnum3dRender,
			Model:      types.V1VideoToVideoCreateBodyStyleModelEnumAbsoluteReality,
			Prompt:     nullable.NewValue("string"),
			PromptType: types.V1VideoToVideoCreateBodyStylePromptTypeEnumAppendDefault,
			Version:    types.V1VideoToVideoCreateBodyStyleVersionEnumDefault,
		},
		Width: nullable.NewValue(512),
	})

	if err != nil {
		t.Fatalf("TestCreate200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
