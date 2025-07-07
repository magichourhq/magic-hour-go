package test_ai_meme_generator_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	ai_meme_generator "github.com/magichourhq/magic-hour-go/resources/v1/ai_meme_generator"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestCreate200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.AiMemeGenerator.Create(ai_meme_generator.CreateRequest{
		Name: nullable.NewValue("My Funny Meme"),
		Style: types.V1AiMemeGeneratorCreateBodyStyle{
			SearchWeb: nullable.NewValue(false),
			Template:  types.V1AiMemeGeneratorCreateBodyStyleTemplateEnumDrakeHotlineBling,
			Topic:     "When the code finally works",
		},
	})

	if err != nil {
		t.Fatalf("TestCreate200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
