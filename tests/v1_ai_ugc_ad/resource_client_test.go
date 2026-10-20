package test_ai_ugc_ad_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	ai_ugc_ad "github.com/magichourhq/magic-hour-go/resources/v1/ai_ugc_ad"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestCreate200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.AiUgcAd.Create(ai_ugc_ad.CreateRequest{
		Assets: types.V1AiUgcAdCreateBodyAssets{
			ImageFilePath: "api-assets/id/image.jpg",
		},
		EndSeconds: 4.0,
		Model:      nullable.NewValue(types.V1AiUgcAdCreateBodyModelEnumKling30),
		Resolution: nullable.NewValue(types.V1AiUgcAdCreateBodyResolutionEnum1080p),
		Script:     "Meet your new everyday essential.",
	})

	if err != nil {
		t.Fatalf("TestCreate200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
