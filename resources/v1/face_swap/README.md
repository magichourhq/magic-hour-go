
### Face Swap video <a name="create"></a>

Create a Face Swap video. The estimated frame cost is calculated using 30 FPS. This amount is deducted from your account balance when a video is queued. Once the video is complete, the cost will be updated based on the actual number of frames rendered.
  
Get more information about this mode at our [product page](https://magichour.ai/products/face-swap).
  

**API Endpoint**: `POST /v1/face-swap`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `assets` | ✓ | Provide the assets for face swap. For video, The `video_source` field determines whether `video_file_path` or `youtube_url` field is used | `V1FaceSwapCreateBodyAssets {FaceMappings: nullable.NewValue([]V1FaceSwapCreateBodyAssetsFaceMappingsItem{V1FaceSwapCreateBodyAssetsFaceMappingsItem {NewFace: "api-assets/id/1234.png",OriginalFace: "api-assets/id/0-0.png",},}),FaceSwapMode: nullable.NewValue(V1FaceSwapCreateBodyAssetsFaceSwapModeEnumAllFaces),ImageFilePath: "image/id/1234.png",VideoFilePath: nullable.NewValue("api-assets/id/1234.mp4"),VideoSource: V1FaceSwapCreateBodyAssetsVideoSourceEnumFile,}` |
| `end_seconds` | ✓ | The end time of the input video in seconds | `15.0` |
| `start_seconds` | ✓ | The start time of the input video in seconds | `0.0` |
| `height` | ✗ | Used to determine the dimensions of the output video.     * If height is provided, width will also be required. The larger value between width and height will be used to determine the maximum output resolution while maintaining the original aspect ratio. * If both height and width are omitted, the video will be resized according to your subscription's maximum resolution, while preserving aspect ratio.  Note: if the video's original resolution is less than the maximum, the video will not be resized.  See our [pricing page](https://magichour.ai/pricing) for more details. | `960` |
| `name` | ✗ | The name of video | `"Face Swap video"` |
| `width` | ✗ | Used to determine the dimensions of the output video.     * If width is provided, height will also be required. The larger value between width and height will be used to determine the maximum output resolution while maintaining the original aspect ratio. * If both height and width are omitted, the video will be resized according to your subscription's maximum resolution, while preserving aspect ratio.  Note: if the video's original resolution is less than the maximum, the video will not be resized.  See our [pricing page](https://magichour.ai/pricing) for more details. | `512` |

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
			ImageFilePath: "image/id/1234.png",
			VideoFilePath: nullable.NewValue("api-assets/id/1234.mp4"),
			VideoSource:   types.V1FaceSwapCreateBodyAssetsVideoSourceEnumFile,
		},
		EndSeconds:   15.0,
		Height:       nullable.NewValue(960),
		Name:         nullable.NewValue("Face Swap video"),
		StartSeconds: 0.0,
		Width:        nullable.NewValue(512),
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
Id: "clx7uu86w0a5qp55yxz315r6r",
}`
