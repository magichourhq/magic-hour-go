package test_image_background_remover_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	image_background_remover "github.com/magichourhq/magic-hour-go/resources/v1/image_background_remover"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestCreate200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.ImageBackgroundRemover.Create(image_background_remover.CreateRequest{
		Assets: types.V1ImageBackgroundRemoverCreateBodyAssets{
			BackgroundImageFilePath: nullable.NewValue("api-assets/id/1234.png"),
			ImageFilePath:           "api-assets/id/1234.png",
		},
		Name: nullable.NewValue("Background Remover image"),
	})

	if err != nil {
		t.Fatalf("TestCreate200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
