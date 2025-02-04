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
	client := sdk.NewClient(sdk.WithBearerAuth("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/magichour/magic-hour/0.8.2"))
	res, err := client.V1.AiHeadshotGenerator.Create(ai_headshot_generator.CreateRequest{Assets: types.PostV1AiHeadshotGeneratorBodyAssets{ImageFilePath: "image/id/1234.png"}})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
