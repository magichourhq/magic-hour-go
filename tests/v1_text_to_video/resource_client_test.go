package test_text_to_video_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	text_to_video "github.com/magichourhq/magic-hour-go/resources/v1/text_to_video"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestCreate200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.TextToVideo.Create(text_to_video.CreateRequest{
		EndSeconds:  5.0,
		Name:        nullable.NewValue("Text To Video video"),
		Orientation: types.V1TextToVideoCreateBodyOrientationEnumLandscape,
		Resolution:  nullable.NewValue(types.V1TextToVideoCreateBodyResolutionEnum1080p),
		Style: types.V1TextToVideoCreateBodyStyle{
			Prompt:      "a dog running",
			QualityMode: nullable.NewValue(types.V1TextToVideoCreateBodyStyleQualityModeEnumQuick),
		},
	})

	if err != nil {
		t.Fatalf("TestCreate200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
