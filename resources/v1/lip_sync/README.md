# v1_lipsync

## Module Functions
### Lip Sync <a name="create"></a>

Create a Lip Sync video. The estimated frame cost is calculated using 30 FPS. This amount is deducted from your account balance when a video is queued. Once the video is complete, the cost will be updated based on the actual number of frames rendered.
  
Get more information about this mode at our [product page](https://magichour.ai/products/lip-sync).
  

**API Endpoint**: `POST /v1/lip-sync`

#### Parameters

| Parameter | Required | Deprecated | Description | Example |
|-----------|:--------:|:----------:|-------------|--------|
| `Assets` | ✓ | ✗ | Provide the assets for lip-sync. For video, The `video_source` field determines whether `video_file_path` or `youtube_url` field is used | `V1LipSyncCreateBodyAssets {AudioFilePath: "api-assets/id/1234.mp3",VideoFilePath: nullable.NewValue("api-assets/id/1234.mp4"),VideoSource: V1LipSyncCreateBodyAssetsVideoSourceEnumFile,}` |
| `└─ AudioFilePath` | ✓ | — | The path of the audio file. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).  Please refer to the [Input File documentation](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) to learn more.  | `"api-assets/id/1234.mp3"` |
| `└─ VideoFilePath` | ✗ | — | Required if `video_source` is `file`. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).  Please refer to the [Input File documentation](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) to learn more.  | `"api-assets/id/1234.mp4"` |
| `└─ VideoSource` | ✓ | — |  | `V1LipSyncCreateBodyAssetsVideoSourceEnumFile` |
| `└─ YoutubeUrl` | ✗ | — | Using a youtube video as the input source. This field is required if `video_source` is `youtube` | `"http://www.example.com"` |
| `EndSeconds` | ✓ | ✗ | The end time of the input video in seconds. This value is used to trim the input video. The value must be greater than 0.1, and more than the start_seconds. | `15.0` |
| `StartSeconds` | ✓ | ✗ | The start time of the input video in seconds. This value is used to trim the input video. The value must be greater than 0. | `0.0` |
| `Height` | ✗ | ✓ | `height` is deprecated and no longer influences the output video's resolution.  Output resolution is determined by the **minimum** of: - The resolution of the input video - The maximum resolution allowed by your subscription tier. See our [pricing page](https://magichour.ai/pricing) for more details.  This field is retained only for backward compatibility and will be removed in a future release. | `nullable.NewValue(123)` |
| `MaxFpsLimit` | ✗ | ✗ | Defines the maximum FPS (frames per second) for the output video. If the input video's FPS is lower than this limit, the output video will retain the input FPS. This is useful for reducing unnecessary frame usage in scenarios where high FPS is not required. | `12.0` |
| `Name` | ✗ | ✗ | The name of video. This value is mainly used for your own identification of the video. | `"Lip Sync video"` |
| `Width` | ✗ | ✓ | `width` is deprecated and no longer influences the output video's resolution.  Output resolution is determined by the **minimum** of: - The resolution of the input video - The maximum resolution allowed by your subscription tier. See our [pricing page](https://magichour.ai/pricing) for more details.  This field is retained only for backward compatibility and will be removed in a future release. | `nullable.NewValue(123)` |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	lip_sync "github.com/magichourhq/magic-hour-go/resources/v1/lip_sync"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.LipSync.Create(lip_sync.CreateRequest{
		Assets: types.V1LipSyncCreateBodyAssets{
			AudioFilePath: "api-assets/id/1234.mp3",
			VideoFilePath: nullable.NewValue("api-assets/id/1234.mp4"),
			VideoSource:   types.V1LipSyncCreateBodyAssetsVideoSourceEnumFile,
		},
		EndSeconds:   15.0,
		MaxFpsLimit:  nullable.NewValue(12.0),
		Name:         nullable.NewValue("Lip Sync video"),
		StartSeconds: 0.0,
	})
}

```

#### Response

##### Type
[V1LipSyncCreateResponse](/types/v1_lip_sync_create_response.go)

##### Example
`V1LipSyncCreateResponse {
CreditsCharged: 450,
EstimatedFrameCost: 450,
Id: "cuid-example",
}`
<!-- CUSTOM DOCS START -->

<!-- CUSTOM DOCS END -->

