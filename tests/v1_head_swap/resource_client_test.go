package test_head_swap_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	head_swap "github.com/magichourhq/magic-hour-go/resources/v1/head_swap"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestCreate200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.HeadSwap.Create(head_swap.CreateRequest{
		Assets: types.V1HeadSwapCreateBodyAssets{
			BodyFilePath: "api-assets/id/1234.png",
			HeadFilePath: "api-assets/id/5678.png",
		},
		MaxResolution: nullable.NewValue(1024),
		Name:          nullable.NewValue("My Head Swap image"),
	})

	if err != nil {
		t.Fatalf("TestCreate200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
