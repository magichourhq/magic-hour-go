package test_video_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	video "github.com/magichourhq/magic-hour-go/resources/v1/estimate/video"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestList200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.Estimate.Video.List(video.ListRequest{
		EndSeconds:        5.0,
		Tool:              types.V1EstimateVideoListToolEnumImageToVideo,
		Audio:             nullable.NewValue(true),
		HasVideoReference: nullable.NewValue(true),
		Model:             nullable.NewValue("kling-3.0"),
		Resolution:        nullable.NewValue("720p"),
		StartSeconds:      nullable.NewValue(123.45),
	})

	if err != nil {
		t.Fatalf("TestList200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}

func TestList200SuccessRequiredOnly(t *testing.T) {
	// Success test using only required fields
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.Estimate.Video.List(video.ListRequest{
		EndSeconds: 5.0,
		Tool:       types.V1EstimateVideoListToolEnumImageToVideo,
	})

	if err != nil {
		t.Fatalf("TestList200SuccessRequiredOnly - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
