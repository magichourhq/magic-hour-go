package test_ai_video_translator_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	ai_video_translator "github.com/magichourhq/magic-hour-go/resources/v1/ai_video_translator"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestCreate200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.AiVideoTranslator.Create(ai_video_translator.CreateRequest{
		Assets: types.V1AiVideoTranslatorCreateBodyAssets{
			VideoFilePath: "api-assets/id/1234.mp4",
		},
		EndSeconds:     15.0,
		Name:           nullable.NewValue("My Video Translator video"),
		Resolution:     nullable.NewValue(types.V1AiVideoTranslatorCreateBodyResolutionEnum720p),
		StartSeconds:   nullable.NewValue(0.0),
		TargetLanguage: types.V1AiVideoTranslatorCreateBodyTargetLanguageEnumSpanish,
	})

	if err != nil {
		t.Fatalf("TestCreate200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
