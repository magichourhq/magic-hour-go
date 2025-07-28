
### Face Swap Photo <a name="create"></a>

Create a face swap photo. Each photo costs 5 credits. The height/width of the output image depends on your subscription. Please refer to our [pricing](/pricing) page for more details

**API Endpoint**: `POST /v1/face-swap-photo`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `assets` | ✓ | Provide the assets for face swap photo | `V1FaceSwapPhotoCreateBodyAssets {FaceMappings: nullable.NewValue([]V1FaceSwapPhotoCreateBodyAssetsFaceMappingsItem{V1FaceSwapPhotoCreateBodyAssetsFaceMappingsItem {NewFace: "api-assets/id/1234.png",OriginalFace: "api-assets/id/0-0.png",},}),FaceSwapMode: nullable.NewValue(V1FaceSwapPhotoCreateBodyAssetsFaceSwapModeEnumAllFaces),SourceFilePath: nullable.NewValue("api-assets/id/1234.png"),TargetFilePath: "api-assets/id/1234.png",}` |
| `name` | ✗ | The name of image | `"Face Swap image"` |

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
			FaceMappings: nullable.NewValue([]types.V1FaceSwapPhotoCreateBodyAssetsFaceMappingsItem{
				types.V1FaceSwapPhotoCreateBodyAssetsFaceMappingsItem{
					NewFace:      "api-assets/id/1234.png",
					OriginalFace: "api-assets/id/0-0.png",
				},
			}),
			FaceSwapMode:   nullable.NewValue(types.V1FaceSwapPhotoCreateBodyAssetsFaceSwapModeEnumAllFaces),
			SourceFilePath: nullable.NewValue("api-assets/id/1234.png"),
			TargetFilePath: "api-assets/id/1234.png",
		},
		Name: nullable.NewValue("Face Swap image"),
	})
}

```

#### Response

##### Type
[V1FaceSwapPhotoCreateResponse](/types/v1_face_swap_photo_create_response.go)

##### Example
`V1FaceSwapPhotoCreateResponse {
CreditsCharged: 5,
FrameCost: 5,
Id: "clx7uu86w0a5qp55yxz315r6r",
}`
