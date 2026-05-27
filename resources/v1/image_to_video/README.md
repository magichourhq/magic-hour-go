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

| Parameter             | Required | Deprecated | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | Example                                                                                                                                   |
| --------------------- | :------: | :--------: | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- |
| `Assets`              |    ✓     |     ✗      | Provide the assets for image-to-video. Sora 2 only supports images with an aspect ratio of `9:16` or `16:9`.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                | `V1ImageToVideoCreateBodyAssets {EndImageFilePath: nullable.NewValue("api-assets/id/1234.png"),ImageFilePath: "api-assets/id/1234.png",}` |
| `└─ EndImageFilePath` |    ✗     |     —      | The image to use as the last frame of the video. * **`ltx-2.3`**: Supports 480p, 720p, 1080p. * **`wan-2.2`**: Not supported * **`kling-2.5`**: Supports 1080p. * **`kling-3.0`**: Supports 720p, 1080p, 4k. * **`veo3.1-lite`**: Not supported * **`veo3.1`**: Not supported * **`seedance`**: Supports 480p, 720p, 1080p. * **`seedance-2.0`**: Supports 480p, 720p. * **`sora-2`**: Not supported                                                                                                                                                                                                                                                                                                                                                                                                        | `"api-assets/id/1234.png"`                                                                                                                |
| `└─ ImageFilePath`    |    ✓     |     —      | The path of the image file. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls). See the [file upload guide](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) for details.                                                                                                                                                                                                                                                                                                                                                                                                                                                   | `"api-assets/id/1234.png"`                                                                                                                |
| `EndSeconds`          |    ✓     |     ✗      | The total duration of the output video in seconds. Supported durations depend on the chosen model: * **`ltx-2.3`**: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 15, 20, 25, 30 * **`wan-2.2`**: 3, 4, 5, 6, 7, 8, 9, 10, 15 * **`kling-2.5`**: 5, 10 * **`kling-3.0`**: 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15 * **`veo3.1-lite`**: 8, 16, 24, 32, 40, 48, 56 * **`veo3.1`**: 4, 6, 8, 16, 24, 32, 40, 48, 56 * **`seedance`**: 4, 5, 6, 7, 8, 9, 10, 11, 12 * **`seedance-2.0`**: 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15 * **`sora-2`**: 4, 8, 12, 24, 36, 48, 60                                                                                                                                                                                                                                                | `5.0`                                                                                                                                     |
| `Audio`               |    ✗     |     ✗      | Whether to include audio in the video. Defaults to `false` if not specified. Audio support varies by model: * **`ltx-2.3`**: Toggle-able: no additional credits for audio * **`wan-2.2`**: Not supported * **`kling-2.5`**: Toggle-able: no additional credits for audio * **`kling-3.0`**: Toggle-able: audio adds extra credits when enabled * **`veo3.1-lite`**: Toggle-able: audio adds extra credits when enabled * **`veo3.1`**: Toggle-able: audio adds extra credits when enabled * **`seedance`**: Not supported * **`seedance-2.0`**: Toggle-able: no additional credits for audio * **`sora-2`**: Toggle-able: no additional credits for audio                                                                                                                                                   | `true`                                                                                                                                    |
| `Height`              |    ✗     |     ✓      | `height` is deprecated and no longer influences the output video's resolution. This field is retained only for backward compatibility and will be removed in a future release.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | `nullable.NewValue(123)`                                                                                                                  |
| `Model`               |    ✗     |     ✗      | The AI model to use for video generation. * `default`: uses our currently recommended model for general use. For paid tiers, defaults to `kling-3.0`. For free tiers, it defaults to `ltx-2.3`. * `ltx-2.3`: Fast iteration with lip-sync & end frame * `wan-2.2`: Fast, strong visuals with effects * `kling-2.5`: Motion, action, and camera control * `kling-3.0`: Cinematic, multi-scene storytelling * `veo3.1-lite`: Fast, affordable, high-quality * `veo3.1`: Realistic visuals and prompt adherence * `seedance`: Fast iteration and start/end frames * `seedance-2.0`: State-of-the-art quality and consistency * `sora-2`: Story-first concepts and creativity If you specify the deprecated model value that includes the `-audio` suffix, this will be the same as included `audio` as `true`. | `V1ImageToVideoCreateBodyModelEnumKling30`                                                                                                |
| `Name`                |    ✗     |     ✗      | Give your video a custom name for easy identification.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | `"My Image To Video video"`                                                                                                               |
| `Resolution`          |    ✗     |     ✗      | Controls the output video resolution. Defaults to `720p` on paid tiers and `480p` on free tiers. * **`ltx-2.3`**: Supports 480p, 720p, 1080p. * **`wan-2.2`**: Supports 480p, 720p, 1080p. * **`kling-2.5`**: Supports 720p, 1080p. * **`kling-3.0`**: Supports 720p, 1080p, 4k. * **`veo3.1-lite`**: Supports 720p, 1080p. * **`veo3.1`**: Supports 720p, 1080p. * **`seedance`**: Supports 480p, 720p, 1080p. * **`seedance-2.0`**: Supports 480p, 720p. * **`sora-2`**: Supports 720p.                                                                                                                                                                                                                                                                                                                   | `V1ImageToVideoCreateBodyResolutionEnum720p`                                                                                              |
| `Style`               |    ✗     |     ✗      | Attributed used to dictate the style of the output                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          | `V1ImageToVideoCreateBodyStyle {Prompt: nullable.NewValue("a dog running"),}`                                                             |
| `└─ HighQuality`      |    ✗     |     ✓      | Deprecated: Please use `resolution` instead. For backward compatibility, * `false` maps to 720p resolution * `true` maps to 1080p resolution This field will be removed in a future version. Use the `resolution` field to directly specify the resolution.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | `true`                                                                                                                                    |
| `└─ Prompt`           |    ✗     |     —      | The prompt used for the video.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | `"a dog running"`                                                                                                                         |
| `└─ QualityMode`      |    ✗     |     ✓      | DEPRECATED: Please use `resolution` field instead. For backward compatibility: * `quick` maps to 720p resolution * `studio` maps to 1080p resolution This field will be removed in a future version. Use the `resolution` field to directly to specify the resolution.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | `V1ImageToVideoCreateBodyStyleQualityModeEnumQuick`                                                                                       |
| `Width`               |    ✗     |     ✓      | `width` is deprecated and no longer influences the output video's resolution. This field is retained only for backward compatibility and will be removed in a future release.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | `nullable.NewValue(123)`                                                                                                                  |

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
			EndImageFilePath: nullable.NewValue("api-assets/id/1234.png"),
			ImageFilePath:    "api-assets/id/1234.png",
		},
		Audio:      nullable.NewValue(true),
		EndSeconds: 5.0,
		Model:      nullable.NewValue(types.V1ImageToVideoCreateBodyModelEnumKling30),
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
