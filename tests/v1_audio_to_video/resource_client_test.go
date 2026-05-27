package test_audio_to_video_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	audio_to_video "github.com/magichourhq/magic-hour-go/resources/v1/audio_to_video"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestCreate200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.AudioToVideo.Create(audio_to_video.CreateRequest{
		Assets: types.V1AudioToVideoCreateBodyAssets{
			AudioFilePath: "api-assets/id/1234.mp3",
			ImageFilePath: nullable.NewValue("api-assets/id/1234.png"),
		},
		EndSeconds:   15.0,
		Name:         nullable.NewValue("My Audio To Video video"),
		Resolution:   nullable.NewValue(types.V1AudioToVideoCreateBodyResolutionEnum720p),
		StartSeconds: nullable.NewValue(0.0),
		Style: nullable.NewValue(types.V1AudioToVideoCreateBodyStyle{
			Prompt: nullable.NewValue("Car driving through a city"),
		}),
	})

	if err != nil {
		t.Fatalf("TestCreate200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
