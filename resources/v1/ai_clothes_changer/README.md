
### AI Clothes Changer <a name="create"></a>

Change outfits in photos in seconds with just a photo reference. Each photo costs 25 credits.

**API Endpoint**: `POST /v1/ai-clothes-changer`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `assets` | ✓ | Provide the assets for clothes changer | `V1AiClothesChangerCreateBodyAssets {GarmentFilePath: "api-assets/id/outfit.png",GarmentType: V1AiClothesChangerCreateBodyAssetsGarmentTypeEnumDresses,PersonFilePath: "api-assets/id/model.png",}` |
| `name` | ✗ | The name of image | `"Clothes Changer image"` |

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

#### Response

##### Type
[V1AiClothesChangerCreateResponse](/types/v1_ai_clothes_changer_create_response.go)

##### Example
`V1AiClothesChangerCreateResponse {
CreditsCharged: 25,
FrameCost: 25,
Id: "clx7uu86w0a5qp55yxz315r6r",
}`
