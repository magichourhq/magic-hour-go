# v1.character_replace

## Module Functions

### Character Replace <a name="create"></a>

**What this API does**

Create the same Character Replace you can make in the browser, but programmatically, so you can automate it, run it at scale, or connect it to your own app or workflow.

**Good for**

- Automation and batch processing
- Adding character replace into apps, pipelines, or tools

**How it works (3 steps)**

1. Upload your inputs (video, image, or audio) with [Generate Upload URLs](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls) and copy the `file_path`.
2. Send a request to create a character replace job with the basic fields.
3. Check the job status until it's `complete`, then download the result from `downloads`.

**Key options**

- Inputs: usually a file, sometimes a YouTube link, depending on project type
- Resolution: free users are limited to 576px; higher plans unlock HD and larger sizes
- Extra fields: e.g. `face_swap_mode`, `start_seconds`/`end_seconds`, or a text prompt

**Cost**\
Credits are only charged for the frames that actually render. You'll see an estimate when the job is queued, and the final total after it's done.

For detailed examples, see the [product page](https://magichour.ai/products/character-replace).

**API Endpoint**: `POST /v1/character-replace`

#### Parameters

| Parameter         | Required | Description                                                                              | Example                                                                                                                                                                                                                                                                                                                                                   |
| ----------------- | :------: | ---------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Data`            |    ✗     |                                                                                          | `V1CharacterReplaceCreateBody {Assets: V1CharacterReplaceCreateBodyAssets {ImageFilePath: "api-assets/id/5678.png",VideoFilePath: "api-assets/id/1234.mp4",},EndSeconds: 15.0,Name: nullable.NewValue("My Character Replace video"),Resolution: nullable.NewValue(V1CharacterReplaceCreateBodyResolutionEnum720p),StartSeconds: nullable.NewValue(0.0),}` |
| `└─ Assets`       |    ✓     | Source video and reference character image for the job.                                  | `V1CharacterReplaceCreateBodyAssets {ImageFilePath: "api-assets/id/5678.png",VideoFilePath: "api-assets/id/1234.mp4",}`                                                                                                                                                                                                                                   |
| `└─ EndSeconds`   |    ✓     | End time of your clip (seconds). Must be greater than start_seconds.                     | `15.0`                                                                                                                                                                                                                                                                                                                                                    |
| `└─ Name`         |    ✗     | Give your video a custom name for easy identification.                                   | `"My Character Replace video"`                                                                                                                                                                                                                                                                                                                            |
| `└─ Resolution`   |    ✗     | Output video resolution. Defaults to 480p, the lowest resolution available on your plan. | `V1CharacterReplaceCreateBodyResolutionEnum720p`                                                                                                                                                                                                                                                                                                          |
| `└─ StartSeconds` |    ✗     | Start time of your clip (seconds). Must be ≥ 0.                                          | `0.0`                                                                                                                                                                                                                                                                                                                                                     |
| `└─ Style`        |    ✗     | Optional style controls for replace vs animate mode and subject selection.               | `V1CharacterReplaceCreateBodyStyle {Mode: nullable.NewValue(V1CharacterReplaceCreateBodyStyleModeEnumReplace),SelectionMode: nullable.NewValue(V1CharacterReplaceCreateBodyStyleSelectionModeEnumAuto),}`                                                                                                                                                 |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	character_replace "github.com/magichourhq/magic-hour-go/resources/v1/character_replace"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.CharacterReplace.Create(character_replace.CreateRequest{})
}
```

#### Response

##### Type

[V1CharacterReplaceCreateResponse](/types/v1_character_replace_create_response.go)

##### Example

```go
V1CharacterReplaceCreateResponse {
CreditsCharged: 450,
EstimatedFrameCost: 450,
Id: "cuid-example",
}
```
