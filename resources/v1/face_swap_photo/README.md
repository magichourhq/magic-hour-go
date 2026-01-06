# v1.face_swap_photo

## Module Functions

### Face Swap Photo <a name="create"></a>

Create a face swap photo. Each photo costs 5 credits. The height/width of the output image depends on your subscription. Please refer to our [pricing](https://magichour.ai/pricing) page for more details

**API Endpoint**: `POST /v1/face-swap-photo`

#### Parameters

| Parameter           | Required | Description                                                                                                                                                                                                                                                                                                                                                                                                                                 | Example                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| ------------------- | :------: | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Assets`            |    ✓     | Provide the assets for face swap photo                                                                                                                                                                                                                                                                                                                                                                                                      | `V1FaceSwapPhotoCreateBodyAssets {FaceMappings: nullable.NewValue([]V1FaceSwapPhotoCreateBodyAssetsFaceMappingsItem{V1FaceSwapPhotoCreateBodyAssetsFaceMappingsItem {NewFace: "api-assets/id/1234.png",OriginalFace: "api-assets/id/0-0.png",},}),FaceSwapMode: nullable.NewValue(V1FaceSwapPhotoCreateBodyAssetsFaceSwapModeEnumAllFaces),SourceFilePath: nullable.NewValue("api-assets/id/1234.png"),TargetFilePath: "api-assets/id/1234.png",}` |
| `└─ FaceMappings`   |    ✗     | This is the array of face mappings used for multiple face swap. The value is required if `face_swap_mode` is `individual-faces`.                                                                                                                                                                                                                                                                                                            | `[]V1FaceSwapPhotoCreateBodyAssetsFaceMappingsItem{V1FaceSwapPhotoCreateBodyAssetsFaceMappingsItem {NewFace: "api-assets/id/1234.png",OriginalFace: "api-assets/id/0-0.png",},}`                                                                                                                                                                                                                                                                   |
| `└─ FaceSwapMode`   |    ✗     | Choose how to swap faces: **all-faces** (recommended) — swap all detected faces using one source image (`source_file_path` required) +- **individual-faces** — specify exact mappings using `face_mappings`                                                                                                                                                                                                                                 | `V1FaceSwapPhotoCreateBodyAssetsFaceSwapModeEnumAllFaces`                                                                                                                                                                                                                                                                                                                                                                                          |
| `└─ SourceFilePath` |    ✗     | This is the image from which the face is extracted. The value is required if `face_swap_mode` is `all-faces`. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls). See the [file upload guide](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) for details. | `"api-assets/id/1234.png"`                                                                                                                                                                                                                                                                                                                                                                                                                         |
| `└─ TargetFilePath` |    ✓     | This is the image where the face from the source image will be placed. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls). See the [file upload guide](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) for details.                                        | `"api-assets/id/1234.png"`                                                                                                                                                                                                                                                                                                                                                                                                                         |
| `Name`              |    ✗     | Give your image a custom name for easy identification.                                                                                                                                                                                                                                                                                                                                                                                      | `"My Face Swap image"`                                                                                                                                                                                                                                                                                                                                                                                                                             |

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
		Name: nullable.NewValue("My Face Swap image"),
	})
}
```

#### Response

##### Type

[V1FaceSwapPhotoCreateResponse](/types/v1_face_swap_photo_create_response.go)

##### Example

```go
V1FaceSwapPhotoCreateResponse {
CreditsCharged: 5,
FrameCost: 5,
Id: "cuid-example",
}
```
