package test_image_background_remover_client

import (
	fmt "fmt"
	sdk "github.com/magichourhq/magic-hour-go/client"
	image_background_remover "github.com/magichourhq/magic-hour-go/resources/v1/image_background_remover"
	types "github.com/magichourhq/magic-hour-go/types"
	testing "testing"
)

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.ImageBackgroundRemover.Create(image_background_remover.CreateRequest{
		Assets: types.PostV1ImageBackgroundRemoverBodyAssets{
			ImageFilePath: "api-assets/id/1234.png",
		},
	})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
