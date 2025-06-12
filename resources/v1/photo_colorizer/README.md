
### Photo Colorizer <a name="create"></a>

Colorize image. Each image costs 5 credits.

**API Endpoint**: `POST /v1/photo-colorizer`

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

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `assets` | ✓ | Provide the assets for photo colorization | `V1PhotoColorizerCreateBodyAssets {ImageFilePath: "api-assets/id/1234.png",}` |
| `name` | ✗ | The name of image | `"Photo Colorizer image"` |
