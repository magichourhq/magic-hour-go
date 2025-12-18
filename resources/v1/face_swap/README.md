# v1.face_swap

## Module Functions

### Face Swap Video <a name="create"></a>

Create a Face Swap video. The estimated frame cost is calculated using 30 FPS. This amount is deducted from your account balance when a video is queued. Once the video is complete, the cost will be updated based on the actual number of frames rendered.
  
Get more information about this mode at our [product page](https://magichour.ai/products/face-swap).
  

**API Endpoint**: `POST /v1/face-swap`

#### Parameters

| Parameter | Required | Deprecated | Description | Example |
|-----------|:--------:|:----------:|-------------|--------|
| `Assets` | ✓ | ✗ | Provide the assets for face swap. For video, The `video_source` field determines whether `video_file_path` or `youtube_url` field is used | `V1FaceSwapCreateBodyAssets {FaceMappings: nullable.NewValue([]V1FaceSwapCreateBodyAssetsFaceMappingsItem{V1FaceSwapCreateBodyAssetsFaceMappingsItem {NewFace: "api-assets/id/1234.png",OriginalFace: "api-assets/id/0-0.png",},}),FaceSwapMode: nullable.NewValue(V1FaceSwapCreateBodyAssetsFaceSwapModeEnumAllFaces),ImageFilePath: nullable.NewValue("image/id/1234.png"),VideoFilePath: nullable.NewValue("api-assets/id/1234.mp4"),VideoSource: V1FaceSwapCreateBodyAssetsVideoSourceEnumFile,}` |
| `└─ FaceMappings` | ✗ | — | This is the array of face mappings used for multiple face swap. The value is required if `face_swap_mode` is `individual-faces`. | `[]V1FaceSwapCreateBodyAssetsFaceMappingsItem{V1FaceSwapCreateBodyAssetsFaceMappingsItem {NewFace: "api-assets/id/1234.png",OriginalFace: "api-assets/id/0-0.png",},}` |
| `└─ FaceSwapMode` | ✗ | — | The mode of face swap. * `all-faces` - Swap all faces in the target image or video. `source_file_path` is required. * `individual-faces` - Swap individual faces in the target image or video. `source_faces` is required. | `V1FaceSwapCreateBodyAssetsFaceSwapModeEnumAllFaces` |
| `└─ ImageFilePath` | ✗ | — | The path of the input image with the face to be swapped.  The value is required if `face_swap_mode` is `all-faces`.  This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).  Please refer to the [Input File documentation](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) to learn more.  | `"image/id/1234.png"` |
| `└─ VideoFilePath` | ✗ | — | Required if `video_source` is `file`. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).  Please refer to the [Input File documentation](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) to learn more.  | `"api-assets/id/1234.mp4"` |
| `└─ VideoSource` | ✓ | — |  | `V1FaceSwapCreateBodyAssetsVideoSourceEnumFile` |
| `└─ YoutubeUrl` | ✗ | — | Using a youtube video as the input source. This field is required if `video_source` is `youtube` | `"http://www.example.com"` |
| `EndSeconds` | ✓ | ✗ | The end time of the input video in seconds. This value is used to trim the input video. The value must be greater than 0.1, and more than the start_seconds. | `15.0` |
| `StartSeconds` | ✓ | ✗ | The start time of the input video in seconds. This value is used to trim the input video. The value must be greater than 0. | `0.0` |
| `Height` | ✗ | ✓ | `height` is deprecated and no longer influences the output video's resolution.  Output resolution is determined by the **minimum** of: - The resolution of the input video - The maximum resolution allowed by your subscription tier. See our [pricing page](https://magichour.ai/pricing) for more details.  This field is retained only for backward compatibility and will be removed in a future release. | `nullable.NewValue(123)` |
| `Name` | ✗ | ✗ | The name of video. This value is mainly used for your own identification of the video. | `"Face Swap video"` |
| `Style` | ✗ | ✗ | Style of the face swap video. | `V1FaceSwapCreateBodyStyle {Version: nullable.NewValue(V1FaceSwapCreateBodyStyleVersionEnumDefault),}` |
| `└─ Version` | ✗ | — | * `v1` - May preserve skin detail and texture better, but weaker identity preservation. * `v2` - Faster, sharper, better handling of hair and glasses. stronger identity preservation. * `default` - Use the version we recommend, which will change over time. This is recommended unless you need a specific earlier version. This is the default behavior. | `V1FaceSwapCreateBodyStyleVersionEnumDefault` |
| `Width` | ✗ | ✓ | `width` is deprecated and no longer influences the output video's resolution.  Output resolution is determined by the **minimum** of: - The resolution of the input video - The maximum resolution allowed by your subscription tier. See our [pricing page](https://magichour.ai/pricing) for more details.  This field is retained only for backward compatibility and will be removed in a future release. | `nullable.NewValue(123)` |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	face_swap "github.com/magichourhq/magic-hour-go/resources/v1/face_swap"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.FaceSwap.Create(face_swap.CreateRequest{
		Assets: types.V1FaceSwapCreateBodyAssets{
			FaceMappings: nullable.NewValue([]types.V1FaceSwapCreateBodyAssetsFaceMappingsItem{
				types.V1FaceSwapCreateBodyAssetsFaceMappingsItem{
					NewFace:      "api-assets/id/1234.png",
					OriginalFace: "api-assets/id/0-0.png",
				},
			}),
			FaceSwapMode:  nullable.NewValue(types.V1FaceSwapCreateBodyAssetsFaceSwapModeEnumAllFaces),
			ImageFilePath: nullable.NewValue("image/id/1234.png"),
			VideoFilePath: nullable.NewValue("api-assets/id/1234.mp4"),
			VideoSource:   types.V1FaceSwapCreateBodyAssetsVideoSourceEnumFile,
		},
		EndSeconds:   15.0,
		Name:         nullable.NewValue("Face Swap video"),
		StartSeconds: 0.0,
		Style: nullable.NewValue(types.V1FaceSwapCreateBodyStyle{
			Version: nullable.NewValue(types.V1FaceSwapCreateBodyStyleVersionEnumDefault),
		}),
	})
}

```

#### Response

##### Type
[V1FaceSwapCreateResponse](/types/v1_face_swap_create_response.go)

##### Example
`V1FaceSwapCreateResponse {
CreditsCharged: 450,
EstimatedFrameCost: 450,
Id: "cuid-example",
}`

