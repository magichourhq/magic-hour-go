package test_ai_image_editor_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	ai_image_editor "github.com/magichourhq/magic-hour-go/resources/v1/ai_image_editor"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestCreate200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.AiImageEditor.Create(ai_image_editor.CreateRequest{
		AspectRatio: nullable.NewValue(types.V1AiImageEditorCreateBodyAspectRatioEnum11),
		Assets: types.V1AiImageEditorCreateBodyAssets{
			ImageFilePath: nullable.NewValue("string"),
			ImageFilePaths: nullable.NewValue([]string{
				"api-assets/id/1234.png",
				"api-assets/id/1235.png",
			}),
		},
		ImageCount: nullable.NewValue(1.0),
		Model:      nullable.NewValue(types.V1AiImageEditorCreateBodyModelEnumDefault),
		Name:       nullable.NewValue("My Ai Image Editor image"),
		Style: types.V1AiImageEditorCreateBodyStyle{
			Model:  nullable.NewValue(types.V1AiImageEditorCreateBodyStyleModelEnumNanoBanana),
			Prompt: "Give me sunglasses",
		},
	})

	if err != nil {
		t.Fatalf("TestCreate200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
