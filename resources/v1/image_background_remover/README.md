
### create <a name="create"></a>
Image Background Remover

Remove background from image. Each image costs 5 frames.

**API Endpoint**: `POST /v1/image-background-remover`

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	image_background_remover "github.com/magichourhq/magic-hour-go/resources/v1/image_background_remover"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.ImageBackgroundRemover.Create(image_background_remover.CreateRequest{
		Assets: types.PostV1ImageBackgroundRemoverBodyAssets{
			ImageFilePath: "api-assets/id/1234.png",
		},
	})
}

```
