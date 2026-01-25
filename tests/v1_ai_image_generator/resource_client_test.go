package test_ai_image_generator_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	ai_image_generator "github.com/magichourhq/magic-hour-go/resources/v1/ai_image_generator"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestCreate200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.AiImageGenerator.Create(ai_image_generator.CreateRequest{
		AspectRatio: nullable.NewValue(types.V1AiImageGeneratorCreateBodyAspectRatioEnum11),
		ImageCount:  1,
		Model:       nullable.NewValue(types.V1AiImageGeneratorCreateBodyModelEnumDefault),
		Name:        nullable.NewValue("My Ai Image image"),
		Orientation: nullable.NewValue(types.V1AiImageGeneratorCreateBodyOrientationEnumLandscape),
		Resolution:  nullable.NewValue(types.V1AiImageGeneratorCreateBodyResolutionEnumAuto),
		Style: types.V1AiImageGeneratorCreateBodyStyle{
			Prompt:      "Cool image",
			QualityMode: nullable.NewValue(types.V1AiImageGeneratorCreateBodyStyleQualityModeEnumStandard),
			Tool:        nullable.NewValue(types.V1AiImageGeneratorCreateBodyStyleToolEnumAiAnimeGenerator),
		},
	})

	if err != nil {
		t.Fatalf("TestCreate200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
