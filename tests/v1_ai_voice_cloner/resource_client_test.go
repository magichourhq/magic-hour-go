package test_ai_voice_cloner_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	ai_voice_cloner "github.com/magichourhq/magic-hour-go/resources/v1/ai_voice_cloner"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestCreate200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.AiVoiceCloner.Create(ai_voice_cloner.CreateRequest{
		Assets: types.V1AiVoiceClonerCreateBodyAssets{
			AudioFilePath: "api-assets/id/1234.mp3",
		},
		Name: nullable.NewValue("Voice Cloner audio"),
		Style: types.V1AiVoiceClonerCreateBodyStyle{
			Prompt: "Hello, this is my cloned voice.",
		},
	})

	if err != nil {
		t.Fatalf("TestCreate200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
