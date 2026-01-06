package test_ai_voice_generator_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	ai_voice_generator "github.com/magichourhq/magic-hour-go/resources/v1/ai_voice_generator"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestCreate200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.AiVoiceGenerator.Create(ai_voice_generator.CreateRequest{
		Name: nullable.NewValue("My Voice Generator audio"),
		Style: types.V1AiVoiceGeneratorCreateBodyStyle{
			Prompt:    "Hello, how are you?",
			VoiceName: types.V1AiVoiceGeneratorCreateBodyStyleVoiceNameEnumElonMusk,
		},
	})

	if err != nil {
		t.Fatalf("TestCreate200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
