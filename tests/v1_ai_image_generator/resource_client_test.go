package test_ai_image_generator_client

import (
	fmt "fmt"
	sdk "github.com/magichourhq/magic-hour-go/client"
	ai_image_generator "github.com/magichourhq/magic-hour-go/resources/v1/ai_image_generator"
	types "github.com/magichourhq/magic-hour-go/types"
	testing "testing"
)

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.AiImageGenerator.Create(ai_image_generator.CreateRequest{
		ImageCount:  1,
		Orientation: types.PostV1AiImageGeneratorBodyOrientationEnumLandscape,
		Style: types.PostV1AiImageGeneratorBodyStyle{
			Prompt: "Cool image",
		},
	})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
