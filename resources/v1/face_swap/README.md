# v1.face_swap

## Module Functions

### Face Swap Video <a name="create"></a>

**What this API does**

Create the same Face Swap you can make in the browser, but programmatically, so you can automate it, run it at scale, or connect it to your own app or workflow.

**Good for**

- Automation and batch processing
- Adding face swap into apps, pipelines, or tools

**How it works (3 steps)**

1. Upload your inputs (video, image, or audio) with [Generate Upload URLs](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls) and copy the `file_path`.
2. Send a request to create a face swap job with the basic fields.
3. Check the job status until it's `complete`, then download the result from `downloads`.

**Key options**

- Inputs: usually a file, sometimes a YouTube link, depending on project type
- Resolution: free users are limited to 576px; higher plans unlock HD and larger sizes
- Extra fields: e.g. `face_swap_mode`, `start_seconds`/`end_seconds`, or a text prompt

**Cost**\
Credits are only charged for the frames that actually render. You'll see an estimate when the job is queued, and the final total after it's done.

For detailed examples, see the [product page](https://magichour.ai/products/face-swap).

**API Endpoint**: `POST /v1/face-swap`

#### Parameters

| Parameter          | Required | Deprecated | Description                                                                                                                                                                                                                                                                                                                                                                                                                                      | Example                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| ------------------ | :------: | :--------: | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Assets`           |    ✓     |     ✗      | Provide the assets for face swap. For video, The `video_source` field determines whether `video_file_path` or `youtube_url` field is used                                                                                                                                                                                                                                                                                                        | `V1FaceSwapCreateBodyAssets {FaceMappings: nullable.NewValue([]V1FaceSwapCreateBodyAssetsFaceMappingsItem{V1FaceSwapCreateBodyAssetsFaceMappingsItem {NewFace: "api-assets/id/1234.png",OriginalFace: "api-assets/id/0-0.png",},}),FaceSwapMode: nullable.NewValue(V1FaceSwapCreateBodyAssetsFaceSwapModeEnumAllFaces),ImageFilePath: nullable.NewValue("image/id/1234.png"),VideoFilePath: nullable.NewValue("api-assets/id/1234.mp4"),VideoSource: V1FaceSwapCreateBodyAssetsVideoSourceEnumFile,}` |
| `└─ FaceMappings`  |    ✗     |     —      | This is the array of face mappings used for multiple face swap. The value is required if `face_swap_mode` is `individual-faces`.                                                                                                                                                                                                                                                                                                                 | `[]V1FaceSwapCreateBodyAssetsFaceMappingsItem{V1FaceSwapCreateBodyAssetsFaceMappingsItem {NewFace: "api-assets/id/1234.png",OriginalFace: "api-assets/id/0-0.png",},}`                                                                                                                                                                                                                                                                                                                                |
| `└─ FaceSwapMode`  |    ✗     |     —      | Choose how to swap faces: **all-faces** (recommended) — swap all detected faces using one source image (`source_file_path` required) +- **individual-faces** — specify exact mappings using `face_mappings`                                                                                                                                                                                                                                      | `V1FaceSwapCreateBodyAssetsFaceSwapModeEnumAllFaces`                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| `└─ ImageFilePath` |    ✗     |     —      | The path of the input image with the face to be swapped. The value is required if `face_swap_mode` is `all-faces`. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls). See the [file upload guide](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) for details. | `"image/id/1234.png"`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| `└─ VideoFilePath` |    ✗     |     —      | Your video file. Required if `video_source` is `file`. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls). See the [file upload guide](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) for details.                                                             | `"api-assets/id/1234.mp4"`                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| `└─ VideoSource`   |    ✓     |     —      | Choose your video source.                                                                                                                                                                                                                                                                                                                                                                                                                        | `V1FaceSwapCreateBodyAssetsVideoSourceEnumFile`                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| `└─ YoutubeUrl`    |    ✗     |     —      | YouTube URL (required if `video_source` is `youtube`).                                                                                                                                                                                                                                                                                                                                                                                           | `"http://www.example.com"`                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| `EndSeconds`       |    ✓     |     ✗      | End time of your clip (seconds). Must be greater than start_seconds.                                                                                                                                                                                                                                                                                                                                                                             | `15.0`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| `StartSeconds`     |    ✓     |     ✗      | Start time of your clip (seconds). Must be ≥ 0.                                                                                                                                                                                                                                                                                                                                                                                                  | `0.0`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| `Height`           |    ✗     |     ✓      | `height` is deprecated and no longer influences the output video's resolution. Output resolution is determined by the **minimum** of: - The resolution of the input video - The maximum resolution allowed by your subscription tier. See our [pricing page](https://magichour.ai/pricing) for more details. This field is retained only for backward compatibility and will be removed in a future release.                                     | `nullable.NewValue(123)`                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| `Name`             |    ✗     |     ✗      | Give your video a custom name for easy identification.                                                                                                                                                                                                                                                                                                                                                                                           | `"My Face Swap video"`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| `Style`            |    ✗     |     ✗      | Style of the face swap video.                                                                                                                                                                                                                                                                                                                                                                                                                    | `V1FaceSwapCreateBodyStyle {Version: nullable.NewValue(V1FaceSwapCreateBodyStyleVersionEnumDefault),}`                                                                                                                                                                                                                                                                                                                                                                                                |
| `└─ Version`       |    ✗     |     —      | * `v1` - May preserve skin detail and texture better, but weaker identity preservation. * `v2` - Faster, sharper, better handling of hair and glasses. stronger identity preservation. * `default` - Use the version we recommend, which will change over time. This is recommended unless you need a specific earlier version. This is the default behavior.                                                                                    | `V1FaceSwapCreateBodyStyleVersionEnumDefault`                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| `Width`            |    ✗     |     ✓      | `width` is deprecated and no longer influences the output video's resolution. Output resolution is determined by the **minimum** of: - The resolution of the input video - The maximum resolution allowed by your subscription tier. See our [pricing page](https://magichour.ai/pricing) for more details. This field is retained only for backward compatibility and will be removed in a future release.                                      | `nullable.NewValue(123)`                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |

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
			ImageFilePath: nullable.NewValue("image/id/1234.png"),
			VideoFilePath: nullable.NewValue("api-assets/id/1234.mp4"),
			VideoSource:   types.V1FaceSwapCreateBodyAssetsVideoSourceEnumFile,
		},
		EndSeconds:   15.0,
		Name:         nullable.NewValue("My Face Swap video"),
		StartSeconds: 0.0,
		Style: nullable.NewValue(types.V1FaceSwapCreateBodyStyle{
			Version: nullable.NewValue(types.V1FaceSwapCreateBodyStyleVersionEnumDefault),
		}),
	})
}
```

#### Response

##### Type

[V1FaceSwapCreateResponse](/types/v1_face_swap_create_response.go)

##### Example

```go
V1FaceSwapCreateResponse {
CreditsCharged: 450,
EstimatedFrameCost: 450,
Id: "cuid-example",
}
```
