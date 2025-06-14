
### Face Swap Photo <a name="create"></a>

Create a face swap photo. Each photo costs 5 credits. The height/width of the output image depends on your subscription. Please refer to our [pricing](/pricing) page for more details

**API Endpoint**: `POST /v1/face-swap-photo`

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	face_swap_photo "github.com/magichourhq/magic-hour-go/resources/v1/face_swap_photo"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.FaceSwapPhoto.Create(face_swap_photo.CreateRequest{
		Assets: types.V1FaceSwapPhotoCreateBodyAssets{
			SourceFilePath: "api-assets/id/1234.png",
			TargetFilePath: "api-assets/id/1234.png",
		},
		Name: nullable.NewValue("Face Swap image"),
	})
}

```

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `assets` | ✓ | Provide the assets for face swap photo | `V1FaceSwapPhotoCreateBodyAssets {SourceFilePath: "api-assets/id/1234.png",TargetFilePath: "api-assets/id/1234.png",}` |
| `name` | ✗ | The name of image | `"Face Swap image"` |
