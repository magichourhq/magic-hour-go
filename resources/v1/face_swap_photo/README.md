# v1.face_swap_photo

## Module Functions

### Face Swap Photo <a name="create"></a>

Create a face swap photo. Each photo costs 5 credits. The height/width of the output image depends on your subscription. Please refer to our [pricing](https://magichour.ai/pricing) page for more details

**API Endpoint**: `POST /v1/face-swap-photo`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `Assets` | ✓ | Provide the assets for face swap photo | `V1FaceSwapPhotoCreateBodyAssets {FaceMappings: nullable.NewValue([]V1FaceSwapPhotoCreateBodyAssetsFaceMappingsItem{V1FaceSwapPhotoCreateBodyAssetsFaceMappingsItem {NewFace: "api-assets/id/1234.png",OriginalFace: "api-assets/id/0-0.png",},}),FaceSwapMode: nullable.NewValue(V1FaceSwapPhotoCreateBodyAssetsFaceSwapModeEnumAllFaces),SourceFilePath: nullable.NewValue("api-assets/id/1234.png"),TargetFilePath: "api-assets/id/1234.png",}` |
| `└─ FaceMappings` | ✗ | This is the array of face mappings used for multiple face swap. The value is required if `face_swap_mode` is `individual-faces`. | `[]V1FaceSwapPhotoCreateBodyAssetsFaceMappingsItem{V1FaceSwapPhotoCreateBodyAssetsFaceMappingsItem {NewFace: "api-assets/id/1234.png",OriginalFace: "api-assets/id/0-0.png",},}` |
| `└─ FaceSwapMode` | ✗ | The mode of face swap. * `all-faces` - Swap all faces in the target image or video. `source_file_path` is required. * `individual-faces` - Swap individual faces in the target image or video. `source_faces` is required. | `V1FaceSwapPhotoCreateBodyAssetsFaceSwapModeEnumAllFaces` |
| `└─ SourceFilePath` | ✗ | This is the image from which the face is extracted. The value is required if `face_swap_mode` is `all-faces`.  This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).  Please refer to the [Input File documentation](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) to learn more.  | `"api-assets/id/1234.png"` |
| `└─ TargetFilePath` | ✓ | This is the image where the face from the source image will be placed. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).  Please refer to the [Input File documentation](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) to learn more.  | `"api-assets/id/1234.png"` |
| `Name` | ✗ | The name of image. This value is mainly used for your own identification of the image. | `"Face Swap image"` |

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
Id: "cuid-example",
}`


