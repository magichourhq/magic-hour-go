package test_lip_sync_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	lip_sync "github.com/magichourhq/magic-hour-go/resources/v1/lip_sync"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithBearerAuth("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/magichour/magic-hour/0.7.2"))
	res, err := client.V1.LipSync.Create(lip_sync.CreateRequest{Assets: types.PostV1LipSyncBodyAssets{AudioFilePath: "audio/id/1234.mp3", VideoFilePath: nullable.NewValue("video/id/1234.mp4"), VideoSource: types.PostV1LipSyncBodyAssetsVideoSourceEnumFile}, EndSeconds: 15, Height: 960, StartSeconds: 0, Width: 512})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
