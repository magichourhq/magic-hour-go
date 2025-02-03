package test_image_to_video_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	image_to_video "github.com/magichourhq/magic-hour-go/resources/v1/image_to_video"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithBearerAuth("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/magichour/magic-hour/0.8.0"))
	res, err := client.V1.ImageToVideo.Create(image_to_video.CreateRequest{Assets: types.PostV1ImageToVideoBodyAssets{ImageFilePath: "image/id/1234.png"}, EndSeconds: 5, Height: 960, Style: types.PostV1ImageToVideoBodyStyle{Prompt: nullable.NewNull[string]()}, Width: 512})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
