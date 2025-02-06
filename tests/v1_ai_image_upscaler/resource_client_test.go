package test_ai_image_upscaler_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	ai_image_upscaler "github.com/magichourhq/magic-hour-go/resources/v1/ai_image_upscaler"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithBearerAuth("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/magichour/magic-hour/0.8.4"))
	res, err := client.V1.AiImageUpscaler.Create(ai_image_upscaler.CreateRequest{Assets: types.PostV1AiImageUpscalerBodyAssets{ImageFilePath: "api-assets/id/1234.png"}, ScaleFactor: 123.45, Style: types.PostV1AiImageUpscalerBodyStyle{Enhancement: types.PostV1AiImageUpscalerBodyStyleEnhancementEnumBalanced}})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
