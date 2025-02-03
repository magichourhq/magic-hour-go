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
	client := sdk.NewClient(sdk.WithBearerAuth("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/magichour/magic-hour/0.8.1"))
	res, err := client.V1.Animation.Create(animation.CreateRequest{Assets: types.PostV1AnimationBodyAssets{AudioFilePath: nullable.NewValue("api-assets/id/1234.mp3"), AudioSource: types.PostV1AnimationBodyAssetsAudioSourceEnumFile, ImageFilePath: nullable.NewValue("api-assets/id/1234.png")}, EndSeconds: 15, Fps: 12, Height: 960, Style: types.PostV1AnimationBodyStyle{ArtStyle: types.PostV1AnimationBodyStyleArtStyleEnumPainterlyIllustration, CameraEffect: types.PostV1AnimationBodyStyleCameraEffectEnumAccelerate, Prompt: nullable.NewValue("Cyberpunk city"), PromptType: types.PostV1AnimationBodyStylePromptTypeEnumAiChoose, TransitionSpeed: 5}, Width: 512})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
