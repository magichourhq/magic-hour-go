package test_ai_clothes_changer_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	ai_clothes_changer "github.com/magichourhq/magic-hour-go/resources/v1/ai_clothes_changer"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestCreate200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithBearerAuth("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/magichour/magic-hour/0.8.2"))
	res, err := client.V1.AiClothesChanger.Create(ai_clothes_changer.CreateRequest{Assets: types.PostV1AiClothesChangerBodyAssets{GarmentFilePath: "api-assets/id/outfit.png", GarmentType: types.PostV1AiClothesChangerBodyAssetsGarmentTypeEnumDresses, PersonFilePath: "api-assets/id/model.png"}})

	if err != nil {
		t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
