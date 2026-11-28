package test_audio_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	audio "github.com/magichourhq/magic-hour-go/resources/v1/estimate/audio"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestList200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.Estimate.Audio.List(audio.ListRequest{
		Tool:           types.V1EstimateAudioListToolEnumAiVoiceGenerator,
		CharacterCount: nullable.NewValue(120),
		Duration:       nullable.NewValue(123),
	})

	if err != nil {
		t.Fatalf("TestList200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}

func TestList200SuccessRequiredOnly(t *testing.T) {
	// Success test using only required fields
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.Estimate.Audio.List(audio.ListRequest{
		Tool: types.V1EstimateAudioListToolEnumAiVoiceGenerator,
	})

	if err != nil {
		t.Fatalf("TestList200SuccessRequiredOnly - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
