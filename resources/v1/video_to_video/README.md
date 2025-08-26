# v1_videotovideo

## Module Functions
### Video-to-Video <a name="create"></a>

Create a Video To Video video. The estimated frame cost is calculated using 30 FPS. This amount is deducted from your account balance when a video is queued. Once the video is complete, the cost will be updated based on the actual number of frames rendered.
  
Get more information about this mode at our [product page](https://magichour.ai/products/video-to-video).
  

**API Endpoint**: `POST /v1/video-to-video`

#### Parameters

| Parameter | Required | Deprecated | Description | Example |
|-----------|:--------:|:----------:|-------------|--------|
| `Assets` | ✓ | ✗ | Provide the assets for video-to-video. For video, The `video_source` field determines whether `video_file_path` or `youtube_url` field is used | `V1VideoToVideoCreateBodyAssets {VideoFilePath: nullable.NewValue("api-assets/id/1234.mp4"),VideoSource: V1VideoToVideoCreateBodyAssetsVideoSourceEnumFile,}` |
| `└─ VideoFilePath` | ✗ | — | Required if `video_source` is `file`. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).  Please refer to the [Input File documentation](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) to learn more.  | `"api-assets/id/1234.mp4"` |
| `└─ VideoSource` | ✓ | — |  | `V1VideoToVideoCreateBodyAssetsVideoSourceEnumFile` |
| `└─ YoutubeUrl` | ✗ | — | Using a youtube video as the input source. This field is required if `video_source` is `youtube` | `"http://www.example.com"` |
| `EndSeconds` | ✓ | ✗ | The end time of the input video in seconds. This value is used to trim the input video. The value must be greater than 0.1, and more than the start_seconds. | `15.0` |
| `StartSeconds` | ✓ | ✗ | The start time of the input video in seconds. This value is used to trim the input video. The value must be greater than 0. | `0.0` |
| `Style` | ✓ | ✗ |  | `V1VideoToVideoCreateBodyStyle {ArtStyle: V1VideoToVideoCreateBodyStyleArtStyleEnum3dRender,Model: V1VideoToVideoCreateBodyStyleModelEnumDefault,Prompt: nullable.NewValue("string"),PromptType: V1VideoToVideoCreateBodyStylePromptTypeEnumDefault,Version: V1VideoToVideoCreateBodyStyleVersionEnumDefault,}` |
| `└─ ArtStyle` | ✓ | — |  | `V1VideoToVideoCreateBodyStyleArtStyleEnum3dRender` |
| `└─ Model` | ✓ | — | * `Dreamshaper` - a good all-around model that works for both animations as well as realism.  * `Absolute Reality` - better at realism, but you'll often get similar results with Dreamshaper as well.  * `Flat 2D Anime` - best for a flat illustration style that's common in most anime. * `default` - use the default recommended model for the selected art style. | `V1VideoToVideoCreateBodyStyleModelEnumDefault` |
| `└─ Prompt` | ✓ | — | The prompt used for the video. Prompt is required if `prompt_type` is `custom` or `append_default`. If `prompt_type` is `default`, then the `prompt` value passed will be ignored. | `nullable.NewValue("string")` |
| `└─ PromptType` | ✓ | — | * `default` - Use the default recommended prompt for the art style. * `custom` - Only use the prompt passed in the API. Note: for v1, lora prompt will still be auto added to apply the art style properly. * `append_default` - Add the default recommended prompt to the end of the prompt passed in the API. | `V1VideoToVideoCreateBodyStylePromptTypeEnumDefault` |
| `└─ Version` | ✓ | — | * `v1` - more detail, closer prompt adherence, and frame-by-frame previews. * `v2` - faster, more consistent, and less noisy. * `default` - use the default version for the selected art style. | `V1VideoToVideoCreateBodyStyleVersionEnumDefault` |
| `FpsResolution` | ✗ | ✗ | Determines whether the resulting video will have the same frame per second as the original video, or half.  * `FULL` - the result video will have the same FPS as the input video * `HALF` - the result video will have half the FPS as the input video | `V1VideoToVideoCreateBodyFpsResolutionEnumHalf` |
| `Height` | ✗ | ✓ | `height` is deprecated and no longer influences the output video's resolution.  Output resolution is determined by the **minimum** of: - The resolution of the input video - The maximum resolution allowed by your subscription tier. See our [pricing page](https://magichour.ai/pricing) for more details.  This field is retained only for backward compatibility and will be removed in a future release. | `nullable.NewValue(123)` |
| `Name` | ✗ | ✗ | The name of video. This value is mainly used for your own identification of the video. | `"Video To Video video"` |
| `Width` | ✗ | ✓ | `width` is deprecated and no longer influences the output video's resolution.  Output resolution is determined by the **minimum** of: - The resolution of the input video - The maximum resolution allowed by your subscription tier. See our [pricing page](https://magichour.ai/pricing) for more details.  This field is retained only for backward compatibility and will be removed in a future release. | `nullable.NewValue(123)` |

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
<!-- CUSTOM DOCS START -->

<!-- CUSTOM DOCS END -->

