package test_lip_sync_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	lip_sync "github.com/magichourhq/magic-hour-go/resources/v1/lip_sync"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestCreate200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.LipSync.Create(lip_sync.CreateRequest{
		Assets: types.V1LipSyncCreateBodyAssets{
			AudioFilePath: "api-assets/id/1234.mp3",
			VideoFilePath: nullable.NewValue("api-assets/id/1234.mp4"),
			VideoSource:   types.V1LipSyncCreateBodyAssetsVideoSourceEnumFile,
			YoutubeUrl:    nullable.NewValue("http://www.example.com"),
		},
		EndSeconds:   15.0,
		Height:       nullable.NewValue(960),
		MaxFpsLimit:  nullable.NewValue(12.0),
		Name:         nullable.NewValue("Lip Sync video"),
		StartSeconds: 0.0,
		Width:        nullable.NewValue(512),
	})

	if err != nil {
		t.Fatalf("TestCreate200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
