# v1_imagetovideo

## Module Functions
### Image-to-Video <a name="create"></a>

Create a Image To Video video. The estimated frame cost is calculated using 30 FPS. This amount is deducted from your account balance when a video is queued. Once the video is complete, the cost will be updated based on the actual number of frames rendered.
  
Get more information about this mode at our [product page](https://magichour.ai/products/image-to-video).
  

**API Endpoint**: `POST /v1/image-to-video`

#### Parameters

| Parameter | Required | Deprecated | Description | Example |
|-----------|:--------:|:----------:|-------------|--------|
| `Assets` | ✓ | ✗ | Provide the assets for image-to-video. | `V1ImageToVideoCreateBodyAssets {ImageFilePath: "api-assets/id/1234.png",}` |
| `└─ ImageFilePath` | ✓ | — | The path of the image file. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).  Please refer to the [Input File documentation](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) to learn more.  | `"api-assets/id/1234.png"` |
| `EndSeconds` | ✓ | ✗ | The total duration of the output video in seconds. | `5.0` |
| `Height` | ✗ | ✓ | `height` is deprecated and no longer influences the output video's resolution.  Output resolution is determined by the **minimum** of: - The resolution of the input video - The maximum resolution allowed by your subscription tier. See our [pricing page](https://magichour.ai/pricing) for more details.  This field is retained only for backward compatibility and will be removed in a future release. | `nullable.NewValue(123)` |
| `Name` | ✗ | ✗ | The name of video. This value is mainly used for your own identification of the video. | `"Image To Video video"` |
| `Resolution` | ✗ | ✗ | Controls the output video resolution. Defaults to `720p` if not specified.  480p and 720p are available on Creator, Pro, or Business tiers. However, 1080p require Pro or Business tier.  **Options:** - `480p` - Supports only 5 or 10 second videos. Output: 24fps. Cost: 120 credits per 5 seconds. - `720p` - Supports videos between 5-60 seconds. Output: 30fps. Cost: 300 credits per 5 seconds. - `1080p` - Supports videos between 5-60 seconds. Output: 30fps. Cost: 600 credits per 5 seconds. | `V1ImageToVideoCreateBodyResolutionEnum720p` |
| `Style` | ✗ | ✗ | Attributed used to dictate the style of the output | `V1ImageToVideoCreateBodyStyle {Prompt: nullable.NewValue("a dog running"),}` |
| `└─ HighQuality` | ✗ | ✓ | Deprecated: Please use `resolution` instead. For backward compatibility,  * `false` maps to 720p resolution * `true` maps to 1080p resolution  This field will be removed in a future version. Use the `resolution` field to directly specify the resolution. | `true` |
| `└─ Prompt` | ✗ | — | The prompt used for the video. | `"a dog running"` |
| `└─ QualityMode` | ✗ | ✓ | DEPRECATED: Please use `resolution` field instead. For backward compatibility: * `quick` maps to 720p resolution * `studio` maps to 1080p resolution  This field will be removed in a future version. Use the `resolution` field to directly to specify the resolution. | `V1ImageToVideoCreateBodyStyleQualityModeEnumQuick` |
| `Width` | ✗ | ✓ | `width` is deprecated and no longer influences the output video's resolution.  Output resolution is determined by the **minimum** of: - The resolution of the input video - The maximum resolution allowed by your subscription tier. See our [pricing page](https://magichour.ai/pricing) for more details.  This field is retained only for backward compatibility and will be removed in a future release. | `nullable.NewValue(123)` |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	image_to_video "github.com/magichourhq/magic-hour-go/resources/v1/image_to_video"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.ImageToVideo.Create(image_to_video.CreateRequest{
		Assets: types.V1ImageToVideoCreateBodyAssets{
			ImageFilePath: "api-assets/id/1234.png",
		},
		EndSeconds: 5.0,
		Name:       nullable.NewValue("Image To Video video"),
		Resolution: nullable.NewValue(types.V1ImageToVideoCreateBodyResolutionEnum720p),
	})
}

```

#### Response

##### Type
[V1ImageToVideoCreateResponse](/types/v1_image_to_video_create_response.go)

##### Example
`V1ImageToVideoCreateResponse {
CreditsCharged: 450,
EstimatedFrameCost: 450,
Id: "cuid-example",
}`
<!-- CUSTOM DOCS START -->

<!-- CUSTOM DOCS END -->

