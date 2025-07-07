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
			EnhanceFace:             false,
			EyeGazeHorizontal:       0.0,
			EyeGazeVertical:         0.0,
			EyeOpenRatio:            0.0,
			EyebrowDirection:        0.0,
			HeadPitch:               0.0,
			HeadRoll:                0.0,
			HeadYaw:                 0.0,
			LipOpenRatio:            0.0,
			MouthGrim:               0.0,
			MouthPositionHorizontal: 0.0,
			MouthPositionVertical:   0.0,
			MouthPout:               0.0,
			MouthPurse:              0.0,
			MouthSmile:              0.0,
		},
	})

	if err != nil {
		t.Fatalf("TestCreate200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
