package test_ai_music_generator_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	ai_music_generator "github.com/magichourhq/magic-hour-go/resources/v1/ai_music_generator"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestCreate200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.AiMusicGenerator.Create(ai_music_generator.CreateRequest{
		Duration: 120,
		Name:     nullable.NewValue("My Music Generator audio"),
		Style: types.V1AiMusicGeneratorCreateBodyStyle{
			Instrumental: nullable.NewValue(false),
			Lyrics:       nullable.NewValue("City lights are fading..."),
			Prompt:       "Dreamy synthwave with warm bass and soft vocals",
		},
	})

	if err != nil {
		t.Fatalf("TestCreate200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
