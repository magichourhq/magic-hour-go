package test_upload_urls_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	upload_urls "github.com/magichourhq/magic-hour-go/resources/v1/files/upload_urls"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.Files.UploadUrls.Create(upload_urls.CreateRequest{
		Items: []types.V1FilesUploadUrlsCreateBodyItemsItem{
			types.V1FilesUploadUrlsCreateBodyItemsItem{
				Extension: "mp4",
				Type:      types.V1FilesUploadUrlsCreateBodyItemsItemTypeEnumVideo,
			},
			types.V1FilesUploadUrlsCreateBodyItemsItem{
				Extension: "mp3",
				Type:      types.V1FilesUploadUrlsCreateBodyItemsItemTypeEnumAudio,
			},
		},
	})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
