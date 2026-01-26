package test_ai_talking_photo_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	ai_talking_photo "github.com/magichourhq/magic-hour-go/resources/v1/ai_talking_photo"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestCreate200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.AiTalkingPhoto.Create(ai_talking_photo.CreateRequest{
		Assets: types.V1AiTalkingPhotoCreateBodyAssets{
			AudioFilePath: "api-assets/id/1234.mp3",
			ImageFilePath: "api-assets/id/1234.png",
		},
		EndSeconds:    15.0,
		MaxResolution: nullable.NewValue(1024),
		Name:          nullable.NewValue("My Talking Photo image"),
		StartSeconds:  0.0,
		Style: nullable.NewValue(types.V1AiTalkingPhotoCreateBodyStyle{
			GenerationMode: nullable.NewValue(types.V1AiTalkingPhotoCreateBodyStyleGenerationModeEnumRealistic),
			Intensity:      nullable.NewValue(123.45),
			Prompt:         nullable.NewValue("string"),
		}),
	})

	if err != nil {
		t.Fatalf("TestCreate200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
