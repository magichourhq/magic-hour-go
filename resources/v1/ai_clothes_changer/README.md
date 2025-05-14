
### AI Clothes Changer <a name="create"></a>

Change outfits in photos in seconds with just a photo reference. Each photo costs 25 credits.

**API Endpoint**: `POST /v1/ai-clothes-changer`

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	ai_clothes_changer "github.com/magichourhq/magic-hour-go/resources/v1/ai_clothes_changer"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.AiClothesChanger.Create(ai_clothes_changer.CreateRequest{
		Assets: types.V1AiClothesChangerCreateBodyAssets{
			GarmentFilePath: "api-assets/id/outfit.png",
			GarmentType:     types.V1AiClothesChangerCreateBodyAssetsGarmentTypeEnumDresses,
			PersonFilePath:  "api-assets/id/model.png",
		},
		Name: nullable.NewValue("Clothes Changer image"),
	})
}

```
