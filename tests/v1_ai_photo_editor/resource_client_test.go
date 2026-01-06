package test_ai_photo_editor_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	ai_photo_editor "github.com/magichourhq/magic-hour-go/resources/v1/ai_photo_editor"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestCreate200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.AiPhotoEditor.Create(ai_photo_editor.CreateRequest{
		Assets: types.V1AiPhotoEditorCreateBodyAssets{
			ImageFilePath: "api-assets/id/1234.png",
		},
		Name:       nullable.NewValue("My Photo Editor image"),
		Resolution: 768,
		Steps:      nullable.NewValue(123),
		Style: types.V1AiPhotoEditorCreateBodyStyle{
			ImageDescription: "A photo of a person",
			LikenessStrength: 5.2,
			NegativePrompt:   nullable.NewValue("painting, cartoon, sketch"),
			Prompt:           "A photo portrait of a person wearing a hat",
			PromptStrength:   3.75,
			Steps:            nullable.NewValue(4),
			UpscaleFactor:    nullable.NewValue(2),
			UpscaleFidelity:  nullable.NewValue(0.5),
		},
	})

	if err != nil {
		t.Fatalf("TestCreate200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
