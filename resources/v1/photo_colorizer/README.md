# v1.photo_colorizer

## Module Functions

### Photo Colorizer <a name="create"></a>

Colorize image. Each image costs 5 credits.

**API Endpoint**: `POST /v1/photo-colorizer`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `Assets` | ✓ | Provide the assets for photo colorization | `V1PhotoColorizerCreateBodyAssets {ImageFilePath: "api-assets/id/1234.png",}` |
| `└─ ImageFilePath` | ✓ | The image used to generate the colorized image. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).  Please refer to the [Input File documentation](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) to learn more.  | `"api-assets/id/1234.png"` |
| `Name` | ✗ | The name of image. This value is mainly used for your own identification of the image. | `"Photo Colorizer image"` |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	photo_colorizer "github.com/magichourhq/magic-hour-go/resources/v1/photo_colorizer"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.PhotoColorizer.Create(photo_colorizer.CreateRequest{
		Assets: types.V1PhotoColorizerCreateBodyAssets{
			ImageFilePath: "api-assets/id/1234.png",
		},
		Name: nullable.NewValue("Photo Colorizer image"),
	})
}

```

#### Response

##### Type
[V1PhotoColorizerCreateResponse](/types/v1_photo_colorizer_create_response.go)

##### Example
`V1PhotoColorizerCreateResponse {
CreditsCharged: 5,
FrameCost: 5,
Id: "cuid-example",
}`

