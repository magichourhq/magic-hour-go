package test_face_swap_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	face_swap "github.com/magichourhq/magic-hour-go/resources/v1/face_swap"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestCreate200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.FaceSwap.Create(face_swap.CreateRequest{
		Assets: types.V1FaceSwapCreateBodyAssets{
			FaceMappings: nullable.NewValue([]types.V1FaceSwapCreateBodyAssetsFaceMappingsItem{
				types.V1FaceSwapCreateBodyAssetsFaceMappingsItem{
					NewFace:      "api-assets/id/1234.png",
					OriginalFace: "api-assets/id/0-0.png",
				},
			}),
			FaceSwapMode:  nullable.NewValue(types.V1FaceSwapCreateBodyAssetsFaceSwapModeEnumAllFaces),
			ImageFilePath: nullable.NewValue("image/id/1234.png"),
			VideoFilePath: nullable.NewValue("api-assets/id/1234.mp4"),
			VideoSource:   types.V1FaceSwapCreateBodyAssetsVideoSourceEnumFile,
			YoutubeUrl:    nullable.NewValue("http://www.example.com"),
		},
		EndSeconds:   15.0,
		Height:       nullable.NewValue(960),
		Name:         nullable.NewValue("Face Swap video"),
		StartSeconds: 0.0,
		Width:        nullable.NewValue(512),
	})

	if err != nil {
		t.Fatalf("TestCreate200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
