package test_character_replace_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	character_replace "github.com/magichourhq/magic-hour-go/resources/v1/character_replace"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestCreate200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.CharacterReplace.Create(character_replace.CreateRequest{
		Data: nullable.NewValue(types.V1CharacterReplaceCreateBody{
			Assets: types.V1CharacterReplaceCreateBodyAssets{
				ImageFilePath: "api-assets/id/5678.png",
				VideoFilePath: "api-assets/id/1234.mp4",
			},
			EndSeconds:   15.0,
			Name:         nullable.NewValue("My Character Replace video"),
			Resolution:   nullable.NewValue(types.V1CharacterReplaceCreateBodyResolutionEnum720p),
			StartSeconds: nullable.NewValue(0.0),
			Style: nullable.NewValue(types.V1CharacterReplaceCreateBodyStyle{
				Mode: nullable.NewValue(types.V1CharacterReplaceCreateBodyStyleModeEnumReplace),
				Points: nullable.NewValue([]types.V1CharacterReplaceCreateBodyStylePointsItem{
					types.V1CharacterReplaceCreateBodyStylePointsItem{
						PositionX:   320,
						PositionY:   180,
						TimeSeconds: 2.5,
					},
				}),
				SelectionMode: nullable.NewValue(types.V1CharacterReplaceCreateBodyStyleSelectionModeEnumAuto),
			}),
		}),
	})

	if err != nil {
		t.Fatalf("TestCreate200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
