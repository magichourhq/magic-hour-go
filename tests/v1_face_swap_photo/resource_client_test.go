package test_face_swap_photo_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	face_swap_photo "github.com/magichourhq/magic-hour-go/resources/v1/face_swap_photo"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestCreate200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.FaceSwapPhoto.Create(face_swap_photo.CreateRequest{
		Assets: types.V1FaceSwapPhotoCreateBodyAssets{
			FaceMappings: nullable.NewValue([]types.V1FaceSwapPhotoCreateBodyAssetsFaceMappingsItem{
				types.V1FaceSwapPhotoCreateBodyAssetsFaceMappingsItem{
					NewFace:      "api-assets/id/1234.png",
					OriginalFace: "api-assets/id/0-0.png",
				},
			}),
			FaceSwapMode:   nullable.NewValue(types.V1FaceSwapPhotoCreateBodyAssetsFaceSwapModeEnumAllFaces),
			SourceFilePath: nullable.NewValue("api-assets/id/1234.png"),
			TargetFilePath: "api-assets/id/1234.png",
		},
		Name: nullable.NewValue("My Face Swap image"),
	})

	if err != nil {
		t.Fatalf("TestCreate200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
