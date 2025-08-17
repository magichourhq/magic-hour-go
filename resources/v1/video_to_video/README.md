
### Video-to-Video <a name="create"></a>

Create a Video To Video video. The estimated frame cost is calculated using 30 FPS. This amount is deducted from your account balance when a video is queued. Once the video is complete, the cost will be updated based on the actual number of frames rendered.
  
Get more information about this mode at our [product page](https://magichour.ai/products/video-to-video).
  

**API Endpoint**: `POST /v1/video-to-video`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `assets` | ✓ | Provide the assets for video-to-video. For video, The `video_source` field determines whether `video_file_path` or `youtube_url` field is used | `V1VideoToVideoCreateBodyAssets {VideoFilePath: nullable.NewValue("api-assets/id/1234.mp4"),VideoSource: V1VideoToVideoCreateBodyAssetsVideoSourceEnumFile,}` |
| `end_seconds` | ✓ | The end time of the input video in seconds. This value is used to trim the input video. The value must be greater than 0.1, and more than the start_seconds. | `15.0` |
| `start_seconds` | ✓ | The start time of the input video in seconds. This value is used to trim the input video. The value must be greater than 0. | `0.0` |
| `style` | ✓ |  | `V1VideoToVideoCreateBodyStyle {ArtStyle: V1VideoToVideoCreateBodyStyleArtStyleEnum3dRender,Model: V1VideoToVideoCreateBodyStyleModelEnumDefault,Prompt: nullable.NewValue("string"),PromptType: V1VideoToVideoCreateBodyStylePromptTypeEnumDefault,Version: V1VideoToVideoCreateBodyStyleVersionEnumDefault,}` |
| `fps_resolution` | ✗ | Determines whether the resulting video will have the same frame per second as the original video, or half.  * `FULL` - the result video will have the same FPS as the input video * `HALF` - the result video will have half the FPS as the input video | `V1VideoToVideoCreateBodyFpsResolutionEnumHalf` |
| `height` | ✗ | `height` is deprecated and no longer influences the output video's resolution.  Output resolution is determined by the **minimum** of: - The resolution of the input video - The maximum resolution allowed by your subscription tier. See our [pricing page](https://magichour.ai/pricing) for more details.  This field is retained only for backward compatibility and will be removed in a future release. | `nullable.NewValue(123)` |
| `name` | ✗ | The name of video. This value is mainly used for your own identification of the video. | `"Video To Video video"` |
| `width` | ✗ | `width` is deprecated and no longer influences the output video's resolution.  Output resolution is determined by the **minimum** of: - The resolution of the input video - The maximum resolution allowed by your subscription tier. See our [pricing page](https://magichour.ai/pricing) for more details.  This field is retained only for backward compatibility and will be removed in a future release. | `nullable.NewValue(123)` |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	video_to_video "github.com/magichourhq/magic-hour-go/resources/v1/video_to_video"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.VideoToVideo.Create(video_to_video.CreateRequest{
		Assets: types.V1VideoToVideoCreateBodyAssets{
			VideoFilePath: nullable.NewValue("api-assets/id/1234.mp4"),
			VideoSource:   types.V1VideoToVideoCreateBodyAssetsVideoSourceEnumFile,
		},
		EndSeconds:    15.0,
		FpsResolution: nullable.NewValue(types.V1VideoToVideoCreateBodyFpsResolutionEnumHalf),
		Name:          nullable.NewValue("Video To Video video"),
		StartSeconds:  0.0,
		Style: types.V1VideoToVideoCreateBodyStyle{
			ArtStyle:   types.V1VideoToVideoCreateBodyStyleArtStyleEnum3dRender,
			Model:      types.V1VideoToVideoCreateBodyStyleModelEnumDefault,
			Prompt:     nullable.NewValue("string"),
			PromptType: types.V1VideoToVideoCreateBodyStylePromptTypeEnumDefault,
			Version:    types.V1VideoToVideoCreateBodyStyleVersionEnumDefault,
		},
	})
}

```

#### Response

##### Type
[V1VideoToVideoCreateResponse](/types/v1_video_to_video_create_response.go)

##### Example
`V1VideoToVideoCreateResponse {
CreditsCharged: 450,
EstimatedFrameCost: 450,
Id: "cuid-example",
}`
