# v1.image_to_video

## Module Functions

### Image-to-Video <a name="create"></a>

**What this API does**

Create the same Image To Video you can make in the browser, but programmatically, so you can automate it, run it at scale, or connect it to your own app or workflow.

**Good for**

- Automation and batch processing
- Adding image to video into apps, pipelines, or tools

**How it works (3 steps)**

1. Upload your inputs (video, image, or audio) with [Generate Upload URLs](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls) and copy the `file_path`.
2. Send a request to create a image to video job with the basic fields.
3. Check the job status until it's `complete`, then download the result from `downloads`.

**Key options**

- Inputs: usually a file, sometimes a YouTube link, depending on project type
- Resolution: free users are limited to 576px; higher plans unlock HD and larger sizes
- Extra fields: e.g. `face_swap_mode`, `start_seconds`/`end_seconds`, or a text prompt

**Cost**\
Credits are only charged for the frames that actually render. You'll see an estimate when the job is queued, and the final total after it's done.

For detailed examples, see the [product page](https://magichour.ai/products/image-to-video).

**API Endpoint**: `POST /v1/image-to-video`

#### Parameters

| Parameter          | Required | Deprecated | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | Example                                                                       |
| ------------------ | :------: | :--------: | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `Assets`           |    ✓     |     ✗      | Provide the assets for image-to-video. Sora 2 only supports images with an aspect ratio of `9:16` or `16:9`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          | `V1ImageToVideoCreateBodyAssets {ImageFilePath: "api-assets/id/1234.png",}`   |
| `└─ ImageFilePath` |    ✓     |     —      | The path of the image file. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls). See the [file upload guide](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) for details.                                                                                                                                                                                                                                             | `"api-assets/id/1234.png"`                                                    |
| `EndSeconds`       |    ✓     |     ✗      | The total duration of the output video in seconds. Supported durations depend on the chosen model: * **Default**: 5-60 seconds (2-12 seconds for 480p). * **Seedance**: 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12 * **Kling 2.5 Audio**: 5, 10 * **Sora 2**: 4, 8, 12, 24, 36, 48, 60 * **Veo 3.1 Audio**: 4, 6, 8, 16, 24, 32, 40, 48, 56 * **Veo 3.1**: 4, 6, 8, 16, 24, 32, 40, 48, 56 * **Kling 1.6**: 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60                                                                                                                                                     | `5.0`                                                                         |
| `Height`           |    ✗     |     ✓      | `height` is deprecated and no longer influences the output video's resolution. Output resolution is determined by the **minimum** of: - The resolution of the input video - The maximum resolution allowed by your subscription tier. See our [pricing page](https://magichour.ai/pricing) for more details. This field is retained only for backward compatibility and will be removed in a future release.                                                                                                                                                                                          | `nullable.NewValue(123)`                                                      |
| `Model`            |    ✗     |     ✗      | The AI model to use for video generation. * `default`: Our recommended model for general use (Kling 2.5 Audio). Note: For backward compatibility, if you use default and end_seconds > 10, we'll fall back to Kling 1.6. * `seedance`: Great for fast iteration and start/end frame * `kling-2.5-audio`: Great for motion, action, and camera control * `sora-2`: Great for story-telling, dialogue & creativity * `veo3.1-audio`: Great for dialogue + SFX generated natively * `veo3.1`: Great for realism, polish, & prompt adherence * `kling-1.6`: Great for dependable clips with smooth motion | `V1ImageToVideoCreateBodyModelEnumKling25Audio`                               |
| `Name`             |    ✗     |     ✗      | Give your video a custom name for easy identification.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                | `"My Image To Video video"`                                                   |
| `Resolution`       |    ✗     |     ✗      | Controls the output video resolution. Defaults to `720p` if not specified. * **Default**: Supports `480p`, `720p`, and `1080p`. * **Seedance**: Supports `480p`, `720p`, `1080p`. * **Kling 2.5 Audio**: Supports `720p`, `1080p`. * **Sora 2**: Supports `720p`. * **Veo 3.1 Audio**: Supports `720p`, `1080p`. * **Veo 3.1**: Supports `720p`, `1080p`. * **Kling 1.6**: Supports `720p`, `1080p`.                                                                                                                                                                                                  | `V1ImageToVideoCreateBodyResolutionEnum720p`                                  |
| `Style`            |    ✗     |     ✗      | Attributed used to dictate the style of the output                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | `V1ImageToVideoCreateBodyStyle {Prompt: nullable.NewValue("a dog running"),}` |
| `└─ HighQuality`   |    ✗     |     ✓      | Deprecated: Please use `resolution` instead. For backward compatibility, * `false` maps to 720p resolution * `true` maps to 1080p resolution This field will be removed in a future version. Use the `resolution` field to directly specify the resolution.                                                                                                                                                                                                                                                                                                                                           | `true`                                                                        |
| `└─ Prompt`        |    ✗     |     —      | The prompt used for the video.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | `"a dog running"`                                                             |
| `└─ QualityMode`   |    ✗     |     ✓      | DEPRECATED: Please use `resolution` field instead. For backward compatibility: * `quick` maps to 720p resolution * `studio` maps to 1080p resolution This field will be removed in a future version. Use the `resolution` field to directly to specify the resolution.                                                                                                                                                                                                                                                                                                                                | `V1ImageToVideoCreateBodyStyleQualityModeEnumQuick`                           |
| `Width`            |    ✗     |     ✓      | `width` is deprecated and no longer influences the output video's resolution. Output resolution is determined by the **minimum** of: - The resolution of the input video - The maximum resolution allowed by your subscription tier. See our [pricing page](https://magichour.ai/pricing) for more details. This field is retained only for backward compatibility and will be removed in a future release.                                                                                                                                                                                           | `nullable.NewValue(123)`                                                      |

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
		Model:      nullable.NewValue(types.V1ImageToVideoCreateBodyModelEnumKling25Audio),
		Name:       nullable.NewValue("My Image To Video video"),
		Resolution: nullable.NewValue(types.V1ImageToVideoCreateBodyResolutionEnum720p),
	})
}
```

#### Response

##### Type

[V1ImageToVideoCreateResponse](/types/v1_image_to_video_create_response.go)

##### Example

```go
V1ImageToVideoCreateResponse {
CreditsCharged: 450,
EstimatedFrameCost: 450,
Id: "cuid-example",
}
```
