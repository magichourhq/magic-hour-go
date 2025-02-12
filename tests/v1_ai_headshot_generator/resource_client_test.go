package test_ai_headshot_generator_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	ai_headshot_generator "github.com/magichourhq/magic-hour-go/resources/v1/ai_headshot_generator"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.AiHeadshotGenerator.Create(ai_headshot_generator.CreateRequest{
		Assets: types.PostV1AiHeadshotGeneratorBodyAssets{
			ImageFilePath: "api-assets/id/1234.png",
		},
	})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
