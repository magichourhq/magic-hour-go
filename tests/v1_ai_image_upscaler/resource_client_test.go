package test_ai_image_upscaler_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	ai_image_upscaler "github.com/magichourhq/magic-hour-go/resources/v1/ai_image_upscaler"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.AiImageUpscaler.Create(ai_image_upscaler.CreateRequest{
		Assets: types.V1AiImageUpscalerCreateBodyAssets{
			ImageFilePath: "api-assets/id/1234.png",
		},
		Name:        nullable.NewValue("Image Upscaler image"),
		ScaleFactor: 2.0,
		Style: types.V1AiImageUpscalerCreateBodyStyle{
			Enhancement: types.V1AiImageUpscalerCreateBodyStyleEnhancementEnumBalanced,
		},
	})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
