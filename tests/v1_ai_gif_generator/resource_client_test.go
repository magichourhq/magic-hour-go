package test_ai_gif_generator_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	ai_gif_generator "github.com/magichourhq/magic-hour-go/resources/v1/ai_gif_generator"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestCreate200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.AiGifGenerator.Create(ai_gif_generator.CreateRequest{
		Name:         nullable.NewValue("My Ai Gif gif"),
		OutputFormat: nullable.NewValue(types.V1AiGifGeneratorCreateBodyOutputFormatEnumGif),
		Style: types.V1AiGifGeneratorCreateBodyStyle{
			Prompt: "Cute dancing cat, pixel art",
		},
	})

	if err != nil {
		t.Fatalf("TestCreate200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
