# v1_aiclotheschanger

## Module Functions
### AI Clothes Changer <a name="create"></a>

Change outfits in photos in seconds with just a photo reference. Each photo costs 25 credits.

**API Endpoint**: `POST /v1/ai-clothes-changer`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `Assets` | ✓ | Provide the assets for clothes changer | `V1AiClothesChangerCreateBodyAssets {GarmentFilePath: "api-assets/id/outfit.png",GarmentType: V1AiClothesChangerCreateBodyAssetsGarmentTypeEnumUpperBody,PersonFilePath: "api-assets/id/model.png",}` |
| `└─ GarmentFilePath` | ✓ | The image of the outfit. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).  Please refer to the [Input File documentation](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) to learn more.  | `"api-assets/id/outfit.png"` |
| `└─ GarmentType` | ✓ | The type of the outfit. | `V1AiClothesChangerCreateBodyAssetsGarmentTypeEnumUpperBody` |
| `└─ PersonFilePath` | ✓ | The image with the person. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).  Please refer to the [Input File documentation](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) to learn more.  | `"api-assets/id/model.png"` |
| `Name` | ✗ | The name of image. This value is mainly used for your own identification of the image. | `"Clothes Changer image"` |

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
			GarmentType:     types.V1AiClothesChangerCreateBodyAssetsGarmentTypeEnumUpperBody,
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
Id: "cuid-example",
}`
<!-- CUSTOM DOCS START -->

<!-- CUSTOM DOCS END -->

