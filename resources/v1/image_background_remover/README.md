# v1.image_background_remover

## Module Functions
### Image Background Remover <a name="create"></a>

Remove background from image. Each image costs 5 credits.

**API Endpoint**: `POST /v1/image-background-remover`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `Assets` | ✓ | Provide the assets for background removal | `V1ImageBackgroundRemoverCreateBodyAssets {BackgroundImageFilePath: nullable.NewValue("api-assets/id/1234.png"),ImageFilePath: "api-assets/id/1234.png",}` |
| `└─ BackgroundImageFilePath` | ✗ | The image used as the new background for the image_file_path. This image will be resized to match the image in image_file_path. Please make sure the resolution between the images are similar.  This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).  Please refer to the [Input File documentation](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) to learn more.  | `"api-assets/id/1234.png"` |
| `└─ ImageFilePath` | ✓ | The image to remove the background. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).  Please refer to the [Input File documentation](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) to learn more.  | `"api-assets/id/1234.png"` |
| `Name` | ✗ | The name of image. This value is mainly used for your own identification of the image. | `"Background Remover image"` |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	image_background_remover "github.com/magichourhq/magic-hour-go/resources/v1/image_background_remover"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.ImageBackgroundRemover.Create(image_background_remover.CreateRequest{
		Assets: types.V1ImageBackgroundRemoverCreateBodyAssets{
			BackgroundImageFilePath: nullable.NewValue("api-assets/id/1234.png"),
			ImageFilePath:           "api-assets/id/1234.png",
		},
		Name: nullable.NewValue("Background Remover image"),
	})
}

```

#### Response

##### Type
[V1ImageBackgroundRemoverCreateResponse](/types/v1_image_background_remover_create_response.go)

##### Example
`V1ImageBackgroundRemoverCreateResponse {
CreditsCharged: 5,
FrameCost: 5,
Id: "cuid-example",
}`
<!-- CUSTOM DOCS START -->

<!-- CUSTOM DOCS END -->

