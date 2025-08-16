package test_ai_face_editor_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	ai_face_editor "github.com/magichourhq/magic-hour-go/resources/v1/ai_face_editor"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestCreate200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.AiFaceEditor.Create(ai_face_editor.CreateRequest{
		Assets: types.V1AiFaceEditorCreateBodyAssets{
			ImageFilePath: "api-assets/id/1234.png",
		},
		Name: nullable.NewValue("Face Editor image"),
		Style: types.V1AiFaceEditorCreateBodyStyle{
			EnhanceFace:             nullable.NewValue(false),
			EyeGazeHorizontal:       nullable.NewValue(0.0),
			EyeGazeVertical:         nullable.NewValue(0.0),
			EyeOpenRatio:            nullable.NewValue(0.0),
			EyebrowDirection:        nullable.NewValue(0.0),
			HeadPitch:               nullable.NewValue(0.0),
			HeadRoll:                nullable.NewValue(0.0),
			HeadYaw:                 nullable.NewValue(0.0),
			LipOpenRatio:            nullable.NewValue(0.0),
			MouthGrim:               nullable.NewValue(0.0),
			MouthPositionHorizontal: nullable.NewValue(0.0),
			MouthPositionVertical:   nullable.NewValue(0.0),
			MouthPout:               nullable.NewValue(0.0),
			MouthPurse:              nullable.NewValue(0.0),
			MouthSmile:              nullable.NewValue(0.0),
		},
	})

	if err != nil {
		t.Fatalf("TestCreate200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
