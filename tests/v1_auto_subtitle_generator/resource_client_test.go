package test_auto_subtitle_generator_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	auto_subtitle_generator "github.com/magichourhq/magic-hour-go/resources/v1/auto_subtitle_generator"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestCreate200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.AutoSubtitleGenerator.Create(auto_subtitle_generator.CreateRequest{
		Assets: types.V1AutoSubtitleGeneratorCreateBodyAssets{
			VideoFilePath: "api-assets/id/1234.mp4",
		},
		EndSeconds:   15.0,
		Name:         nullable.NewValue("My Auto Subtitle video"),
		StartSeconds: 0.0,
		Style: types.V1AutoSubtitleGeneratorCreateBodyStyle{
			CustomConfig: nullable.NewValue(types.V1AutoSubtitleGeneratorCreateBodyStyleCustomConfig{
				Font:                 nullable.NewValue("Noto Sans"),
				FontSize:             nullable.NewValue(24.0),
				FontStyle:            nullable.NewValue("normal"),
				HighlightedTextColor: nullable.NewValue("#FFD700"),
				HorizontalPosition:   nullable.NewValue("center"),
				StrokeColor:          nullable.NewValue("#000000"),
				StrokeWidth:          nullable.NewValue(1.0),
				TextColor:            nullable.NewValue("#FFFFFF"),
				VerticalPosition:     nullable.NewValue("bottom"),
			}),
			Template: nullable.NewValue(types.V1AutoSubtitleGeneratorCreateBodyStyleTemplateEnumCinematic),
		},
	})

	if err != nil {
		t.Fatalf("TestCreate200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
