package test_ai_photo_editor_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	ai_photo_editor "github.com/magichourhq/magic-hour-go/resources/v1/ai_photo_editor"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithBearerAuth("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/magichour/magic-hour/0.8.2"))
	res, err := client.V1.AiPhotoEditor.Create(ai_photo_editor.CreateRequest{Assets: types.PostV1AiPhotoEditorBodyAssets{ImageFilePath: "api-assets/id/1234.png"}, Resolution: 768, Style: types.PostV1AiPhotoEditorBodyStyle{ImageDescription: "A photo of a person", LikenessStrength: 5.2, NegativePrompt: nullable.NewValue("painting, cartoon, sketch"), Prompt: "A photo portrait of a person wearing a hat", PromptStrength: 3.75, Steps: nullable.NewValue(4)}})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
