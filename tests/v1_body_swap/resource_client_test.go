package test_body_swap_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	body_swap "github.com/magichourhq/magic-hour-go/resources/v1/body_swap"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestCreate200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.BodySwap.Create(body_swap.CreateRequest{
		Assets: types.V1BodySwapCreateBodyAssets{
			PersonFilePath: "api-assets/id/1234.png",
			SceneFilePath:  "api-assets/id/5678.png",
		},
		Name:       nullable.NewValue("My Body Swap image"),
		Resolution: types.V1BodySwapCreateBodyResolutionEnum1k,
	})

	if err != nil {
		t.Fatalf("TestCreate200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
