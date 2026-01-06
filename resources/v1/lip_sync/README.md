# v1.lip_sync

## Module Functions

### Lip Sync <a name="create"></a>

**What this API does**

Create the same Lip Sync you can make in the browser, but programmatically, so you can automate it, run it at scale, or connect it to your own app or workflow.

**Good for**

- Automation and batch processing
- Adding lip sync into apps, pipelines, or tools

**How it works (3 steps)**

1. Upload your inputs (video, image, or audio) with [Generate Upload URLs](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls) and copy the `file_path`.
2. Send a request to create a lip sync job with the basic fields.
3. Check the job status until it's `complete`, then download the result from `downloads`.

**Key options**

- Inputs: usually a file, sometimes a YouTube link, depending on project type
- Resolution: free users are limited to 512px; higher plans unlock HD and larger sizes
- Extra fields: e.g. `face_swap_mode`, `start_seconds`/`end_seconds`, or a text prompt

**Cost**\
Credits are only charged for the frames that actually render. You'll see an estimate when the job is queued, and the final total after it's done.

For detailed examples, see the [product page](https://magichour.ai/products/lip-sync).

**API Endpoint**: `POST /v1/lip-sync`

#### Parameters

| Parameter           | Required | Deprecated | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | Example                                                                                                                                                                                     |
| ------------------- | :------: | :--------: | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Assets`            |    ✓     |     ✗      | Provide the assets for lip-sync. For video, The `video_source` field determines whether `video_file_path` or `youtube_url` field is used                                                                                                                                                                                                                                                                                                                                                         | `V1LipSyncCreateBodyAssets {AudioFilePath: "api-assets/id/1234.mp3",VideoFilePath: nullable.NewValue("api-assets/id/1234.mp4"),VideoSource: V1LipSyncCreateBodyAssetsVideoSourceEnumFile,}` |
| `└─ AudioFilePath`  |    ✓     |     —      | The path of the audio file. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls). See the [file upload guide](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) for details.                                                                                                                                        | `"api-assets/id/1234.mp3"`                                                                                                                                                                  |
| `└─ VideoFilePath`  |    ✗     |     —      | Your video file. Required if `video_source` is `file`. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls). See the [file upload guide](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) for details.                                                                                                             | `"api-assets/id/1234.mp4"`                                                                                                                                                                  |
| `└─ VideoSource`    |    ✓     |     —      | Choose your video source.                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | `V1LipSyncCreateBodyAssetsVideoSourceEnumFile`                                                                                                                                              |
| `└─ YoutubeUrl`     |    ✗     |     —      | YouTube URL (required if `video_source` is `youtube`).                                                                                                                                                                                                                                                                                                                                                                                                                                           | `"http://www.example.com"`                                                                                                                                                                  |
| `EndSeconds`        |    ✓     |     ✗      | End time of your clip (seconds). Must be greater than start_seconds.                                                                                                                                                                                                                                                                                                                                                                                                                             | `15.0`                                                                                                                                                                                      |
| `StartSeconds`      |    ✓     |     ✗      | Start time of your clip (seconds). Must be ≥ 0.                                                                                                                                                                                                                                                                                                                                                                                                                                                  | `0.0`                                                                                                                                                                                       |
| `Height`            |    ✗     |     ✓      | `height` is deprecated and no longer influences the output video's resolution. Output resolution is determined by the **minimum** of: - The resolution of the input video - The maximum resolution allowed by your subscription tier. See our [pricing page](https://magichour.ai/pricing) for more details. This field is retained only for backward compatibility and will be removed in a future release.                                                                                     | `nullable.NewValue(123)`                                                                                                                                                                    |
| `MaxFpsLimit`       |    ✗     |     ✗      | Defines the maximum FPS (frames per second) for the output video. If the input video's FPS is lower than this limit, the output video will retain the input FPS. This is useful for reducing unnecessary frame usage in scenarios where high FPS is not required.                                                                                                                                                                                                                                | `12.0`                                                                                                                                                                                      |
| `Name`              |    ✗     |     ✗      | Give your video a custom name for easy identification.                                                                                                                                                                                                                                                                                                                                                                                                                                           | `"My Lip Sync video"`                                                                                                                                                                       |
| `Style`             |    ✗     |     ✗      | Attributes used to dictate the style of the output                                                                                                                                                                                                                                                                                                                                                                                                                                               | `V1LipSyncCreateBodyStyle {GenerationMode: nullable.NewValue(V1LipSyncCreateBodyStyleGenerationModeEnumLite),}`                                                                             |
| `└─ GenerationMode` |    ✗     |     —      | A specific version of our lip sync system, optimized for different needs. * `lite` - Fast and affordable lip sync - best for simple videos. Costs 1 credit per frame of video. * `standard` - Natural, accurate lip sync - best for most creators. Costs 1 credit per frame of video. * `pro` - Premium fidelity with enhanced detail - best for professionals. Costs 2 credits per frame of video. Note: `standard` and `pro` are only available for users on Creator, Pro, and Business tiers. | `V1LipSyncCreateBodyStyleGenerationModeEnumLite`                                                                                                                                            |
| `Width`             |    ✗     |     ✓      | `width` is deprecated and no longer influences the output video's resolution. Output resolution is determined by the **minimum** of: - The resolution of the input video - The maximum resolution allowed by your subscription tier. See our [pricing page](https://magichour.ai/pricing) for more details. This field is retained only for backward compatibility and will be removed in a future release.                                                                                      | `nullable.NewValue(123)`                                                                                                                                                                    |

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
		Name:         nullable.NewValue("My Lip Sync video"),
		StartSeconds: 0.0,
	})
}
```

#### Response

##### Type

[V1LipSyncCreateResponse](/types/v1_lip_sync_create_response.go)

##### Example

```go
V1LipSyncCreateResponse {
CreditsCharged: 450,
EstimatedFrameCost: 450,
Id: "cuid-example",
}
```
