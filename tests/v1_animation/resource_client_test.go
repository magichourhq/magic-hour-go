package test_animation_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	animation "github.com/magichourhq/magic-hour-go/resources/v1/animation"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.Animation.Create(animation.CreateRequest{
		Assets: types.V1AnimationCreateBodyAssets{
			AudioFilePath: nullable.NewValue("api-assets/id/1234.mp3"),
			AudioSource:   types.V1AnimationCreateBodyAssetsAudioSourceEnumFile,
			ImageFilePath: nullable.NewValue("api-assets/id/1234.png"),
		},
		EndSeconds: 15.0,
		Fps:        12.0,
		Height:     960,
		Style: types.V1AnimationCreateBodyStyle{
			ArtStyle:        types.V1AnimationCreateBodyStyleArtStyleEnumPainterlyIllustration,
			CameraEffect:    types.V1AnimationCreateBodyStyleCameraEffectEnumAccelerate,
			Prompt:          nullable.NewValue("Cyberpunk city"),
			PromptType:      types.V1AnimationCreateBodyStylePromptTypeEnumAiChoose,
			TransitionSpeed: 5,
		},
		Width: 512,
	})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
