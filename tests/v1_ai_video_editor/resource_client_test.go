package test_ai_video_editor_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	ai_video_editor "github.com/magichourhq/magic-hour-go/resources/v1/ai_video_editor"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestCreate200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.AiVideoEditor.Create(ai_video_editor.CreateRequest{
		Assets: types.V1AiVideoEditorCreateBodyAssets{
			VideoFilePath: "api-assets/id/1234.mp4",
		},
		EndSeconds:   5.0,
		Name:         nullable.NewValue("My Video Editor video"),
		StartSeconds: nullable.NewValue(0.0),
		Style: types.V1AiVideoEditorCreateBodyStyle{
			Prompt: "Change the car color to blue",
		},
	})

	if err != nil {
		t.Fatalf("TestCreate200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
