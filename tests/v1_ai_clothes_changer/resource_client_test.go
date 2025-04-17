package test_ai_clothes_changer_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	ai_clothes_changer "github.com/magichourhq/magic-hour-go/resources/v1/ai_clothes_changer"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.AiClothesChanger.Create(ai_clothes_changer.CreateRequest{
		Assets: types.V1AiClothesChangerCreateBodyAssets{
			GarmentFilePath: "api-assets/id/outfit.png",
			GarmentType:     types.V1AiClothesChangerCreateBodyAssetsGarmentTypeEnumDresses,
			PersonFilePath:  "api-assets/id/model.png",
		},
		Name: nullable.NewValue("Clothes Changer image"),
	})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
