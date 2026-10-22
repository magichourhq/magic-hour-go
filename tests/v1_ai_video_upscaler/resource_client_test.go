package test_ai_video_upscaler_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	ai_video_upscaler "github.com/magichourhq/magic-hour-go/resources/v1/ai_video_upscaler"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestCreate200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.AiVideoUpscaler.Create(ai_video_upscaler.CreateRequest{
		Assets: types.V1AiVideoUpscalerCreateBodyAssets{
			VideoFilePath: "api-assets/id/1234.mp4",
		},
		EndSeconds:   5.0,
		Name:         nullable.NewValue("My Video Upscaler video"),
		Resolution:   types.V1AiVideoUpscalerCreateBodyResolutionEnum1080p,
		StartSeconds: nullable.NewValue(0.0),
	})

	if err != nil {
		t.Fatalf("TestCreate200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
