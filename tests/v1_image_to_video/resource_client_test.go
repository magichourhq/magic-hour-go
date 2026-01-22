package test_image_to_video_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	image_to_video "github.com/magichourhq/magic-hour-go/resources/v1/image_to_video"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestCreate200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.ImageToVideo.Create(image_to_video.CreateRequest{
		Assets: types.V1ImageToVideoCreateBodyAssets{
			ImageFilePath: "api-assets/id/1234.png",
		},
		EndSeconds: 5.0,
		Height:     nullable.NewValue(123),
		Model:      nullable.NewValue(types.V1ImageToVideoCreateBodyModelEnumSora2),
		Name:       nullable.NewValue("My Image To Video video"),
		Resolution: nullable.NewValue(types.V1ImageToVideoCreateBodyResolutionEnum720p),
		Style: nullable.NewValue(types.V1ImageToVideoCreateBodyStyle{
			HighQuality: nullable.NewValue(true),
			Prompt:      nullable.NewValue("a dog running"),
			QualityMode: nullable.NewValue(types.V1ImageToVideoCreateBodyStyleQualityModeEnumQuick),
		}),
		Width: nullable.NewValue(123),
	})

	if err != nil {
		t.Fatalf("TestCreate200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
