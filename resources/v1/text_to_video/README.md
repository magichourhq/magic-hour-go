# v1.text_to_video

## Module Functions

### Text-to-Video <a name="create"></a>

**What this API does**

Create the same Text To Video you can make in the browser, but programmatically, so you can automate it, run it at scale, or connect it to your own app or workflow.

**Good for**

- Automation and batch processing
- Adding text to video into apps, pipelines, or tools

**How it works (3 steps)**

1. Upload your inputs (video, image, or audio) with [Generate Upload URLs](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls) and copy the `file_path`.
2. Send a request to create a text to video job with the basic fields.
3. Check the job status until it's `complete`, then download the result from `downloads`.

**Key options**

- Inputs: usually a file, sometimes a YouTube link, depending on project type
- Resolution: free users are limited to 576px; higher plans unlock HD and larger sizes
- Extra fields: e.g. `face_swap_mode`, `start_seconds`/`end_seconds`, or a text prompt

**Cost**\
Credits are only charged for the frames that actually render. You'll see an estimate when the job is queued, and the final total after it's done.

For detailed examples, see the [product page](https://magichour.ai/products/text-to-video).

**API Endpoint**: `POST /v1/text-to-video`

#### Parameters

| Parameter        | Required | Deprecated | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | Example                                                   |
| ---------------- | :------: | :--------: | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------- |
| `EndSeconds`     |    ✓     |     ✗      | The total duration of the output video in seconds. Supported durations depend on the chosen model: * **Default**: 5-60 seconds (2-12 seconds for 480p). * **Seedance**: 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12 * **Kling 2.5 Audio**: 5, 10 * **Sora 2**: 4, 8, 12, 24, 36, 48, 60 * **Veo 3.1 Audio**: 4, 6, 8, 16, 24, 32, 40, 48, 56 * **Veo 3.1**: 4, 6, 8, 16, 24, 32, 40, 48, 56 * **Kling 1.6**: 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60                                                                                                                                                     | `5.0`                                                     |
| `Style`          |    ✓     |     ✗      |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | `V1TextToVideoCreateBodyStyle {Prompt: "a dog running",}` |
| `└─ Prompt`      |    ✓     |     —      | The prompt used for the video.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | `"a dog running"`                                         |
| `└─ QualityMode` |    ✗     |     ✓      | DEPRECATED: Please use `resolution` field instead. For backward compatibility: * `quick` maps to 720p resolution * `studio` maps to 1080p resolution This field will be removed in a future version. Use the `resolution` field to directly to specify the resolution.                                                                                                                                                                                                                                                                                                                                | `V1TextToVideoCreateBodyStyleQualityModeEnumQuick`        |
| `AspectRatio`    |    ✗     |     ✗      | Determines the aspect ratio of the output video. * **Seedance**: Supports `9:16`, `16:9`, `1:1`. * **Kling 2.5 Audio**: Supports `9:16`, `16:9`, `1:1`. * **Sora 2**: Supports `9:16`, `16:9`. * **Veo 3.1 Audio**: Supports `9:16`, `16:9`. * **Veo 3.1**: Supports `9:16`, `16:9`. * **Kling 1.6**: Supports `9:16`, `16:9`, `1:1`.                                                                                                                                                                                                                                                                 | `V1TextToVideoCreateBodyAspectRatioEnum169`               |
| `Model`          |    ✗     |     ✗      | The AI model to use for video generation. * `default`: Our recommended model for general use (Kling 2.5 Audio). Note: For backward compatibility, if you use default and end_seconds > 10, we'll fall back to Kling 1.6. * `seedance`: Great for fast iteration and start/end frame * `kling-2.5-audio`: Great for motion, action, and camera control * `sora-2`: Great for story-telling, dialogue & creativity * `veo3.1-audio`: Great for dialogue + SFX generated natively * `veo3.1`: Great for realism, polish, & prompt adherence * `kling-1.6`: Great for dependable clips with smooth motion | `V1TextToVideoCreateBodyModelEnumKling25Audio`            |
| `Name`           |    ✗     |     ✗      | Give your video a custom name for easy identification.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                | `"My Text To Video video"`                                |
| `Orientation`    |    ✗     |     ✓      | Deprecated. Use `aspect_ratio` instead.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | `V1TextToVideoCreateBodyOrientationEnumLandscape`         |
| `Resolution`     |    ✗     |     ✗      | Controls the output video resolution. Defaults to `720p` if not specified. * **Default**: Supports `480p`, `720p`, and `1080p`. * **Seedance**: Supports `480p`, `720p`, `1080p`. * **Kling 2.5 Audio**: Supports `720p`, `1080p`. * **Sora 2**: Supports `720p`. * **Veo 3.1 Audio**: Supports `720p`, `1080p`. * **Veo 3.1**: Supports `720p`, `1080p`. * **Kling 1.6**: Supports `720p`, `1080p`.                                                                                                                                                                                                  | `V1TextToVideoCreateBodyResolutionEnum720p`               |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	text_to_video "github.com/magichourhq/magic-hour-go/resources/v1/text_to_video"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.TextToVideo.Create(text_to_video.CreateRequest{
		AspectRatio: nullable.NewValue(types.V1TextToVideoCreateBodyAspectRatioEnum169),
		EndSeconds:  5.0,
		Model:       nullable.NewValue(types.V1TextToVideoCreateBodyModelEnumKling25Audio),
		Name:        nullable.NewValue("My Text To Video video"),
		Orientation: nullable.NewValue(types.V1TextToVideoCreateBodyOrientationEnumLandscape),
		Resolution:  nullable.NewValue(types.V1TextToVideoCreateBodyResolutionEnum720p),
		Style: types.V1TextToVideoCreateBodyStyle{
			Prompt: "a dog running",
		},
	})
}
```

#### Response

##### Type

[V1TextToVideoCreateResponse](/types/v1_text_to_video_create_response.go)

##### Example

```go
V1TextToVideoCreateResponse {
CreditsCharged: 450,
EstimatedFrameCost: 450,
Id: "cuid-example",
}
```
