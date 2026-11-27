package test_image_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	image "github.com/magichourhq/magic-hour-go/resources/v1/estimate/image"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestList200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.Estimate.Image.List(image.ListRequest{
		Tool:        types.V1EstimateImageListToolEnumAiImageEditor,
		ImageCount:  nullable.NewValue(4),
		Mode:        nullable.NewValue("string"),
		Model:       nullable.NewValue("nano-banana-2.1"),
		Resolution:  nullable.NewValue("2k"),
		ScaleFactor: nullable.NewValue(123.45),
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
	res, err := client.V1.Estimate.Image.List(image.ListRequest{
		Tool: types.V1EstimateImageListToolEnumAiImageEditor,
	})

	if err != nil {
		t.Fatalf("TestList200SuccessRequiredOnly - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
