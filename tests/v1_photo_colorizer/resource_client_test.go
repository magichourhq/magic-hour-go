package test_photo_colorizer_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	photo_colorizer "github.com/magichourhq/magic-hour-go/resources/v1/photo_colorizer"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.PhotoColorizer.Create(photo_colorizer.CreateRequest{
		Assets: types.V1PhotoColorizerCreateBodyAssets{
			ImageFilePath: "api-assets/id/1234.png",
		},
		Name: nullable.NewValue("Photo Colorizer image"),
	})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
