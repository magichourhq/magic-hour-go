
### Video-to-Video <a name="create"></a>

Create a Video To Video video. The estimated frame cost is calculated using 30 FPS. This amount is deducted from your account balance when a video is queued. Once the video is complete, the cost will be updated based on the actual number of frames rendered.
  
Get more information about this mode at our [product page](/products/video-to-video).
  

**API Endpoint**: `POST /v1/video-to-video`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `assets` | ✓ | Provide the assets for video-to-video. For video, The `video_source` field determines whether `video_file_path` or `youtube_url` field is used | `V1VideoToVideoCreateBodyAssets {VideoFilePath: nullable.NewValue("api-assets/id/1234.mp4"),VideoSource: V1VideoToVideoCreateBodyAssetsVideoSourceEnumFile,}` |
| `end_seconds` | ✓ | The end time of the input video in seconds | `15.0` |
| `start_seconds` | ✓ | The start time of the input video in seconds | `0.0` |
| `style` | ✓ |  | `V1VideoToVideoCreateBodyStyle {ArtStyle: V1VideoToVideoCreateBodyStyleArtStyleEnum3dRender,Model: V1VideoToVideoCreateBodyStyleModelEnumAbsoluteReality,Prompt: nullable.NewValue("string"),PromptType: V1VideoToVideoCreateBodyStylePromptTypeEnumAppendDefault,Version: V1VideoToVideoCreateBodyStyleVersionEnumDefault,}` |
| `fps_resolution` | ✗ | Determines whether the resulting video will have the same frame per second as the original video, or half.  * `FULL` - the result video will have the same FPS as the input video * `HALF` - the result video will have half the FPS as the input video | `V1VideoToVideoCreateBodyFpsResolutionEnumHalf` |
| `height` | ✗ | Used to determine the dimensions of the output video.     * If height is provided, width will also be required. The larger value between width and height will be used to determine the maximum output resolution while maintaining the original aspect ratio. * If both height and width are omitted, the video will be resized according to your subscription's maximum resolution, while preserving aspect ratio.  Note: if the video's original resolution is less than the maximum, the video will not be resized.  See our [pricing page](https://magichour.ai/pricing) for more details. | `960` |
| `name` | ✗ | The name of video | `"Video To Video video"` |
| `width` | ✗ | Used to determine the dimensions of the output video.     * If width is provided, height will also be required. The larger value between width and height will be used to determine the maximum output resolution while maintaining the original aspect ratio. * If both height and width are omitted, the video will be resized according to your subscription's maximum resolution, while preserving aspect ratio.  Note: if the video's original resolution is less than the maximum, the video will not be resized.  See our [pricing page](https://magichour.ai/pricing) for more details. | `512` |

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
		Height:        nullable.NewValue(960),
		Name:          nullable.NewValue("Video To Video video"),
		StartSeconds:  0.0,
		Style: types.V1VideoToVideoCreateBodyStyle{
			ArtStyle:   types.V1VideoToVideoCreateBodyStyleArtStyleEnum3dRender,
			Model:      types.V1VideoToVideoCreateBodyStyleModelEnumAbsoluteReality,
			Prompt:     nullable.NewValue("string"),
			PromptType: types.V1VideoToVideoCreateBodyStylePromptTypeEnumAppendDefault,
			Version:    types.V1VideoToVideoCreateBodyStyleVersionEnumDefault,
		},
		Width: nullable.NewValue(512),
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
Id: "clx7uu86w0a5qp55yxz315r6r",
}`
