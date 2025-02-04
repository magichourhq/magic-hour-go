package test_text_to_video_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	text_to_video "github.com/magichourhq/magic-hour-go/resources/v1/text_to_video"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithBearerAuth("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/magichour/magic-hour/0.8.2"))
	res, err := client.V1.TextToVideo.Create(text_to_video.CreateRequest{EndSeconds: 5.0, Orientation: types.PostV1TextToVideoBodyOrientationEnumLandscape, Style: types.PostV1TextToVideoBodyStyle{Prompt: "string"}})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
