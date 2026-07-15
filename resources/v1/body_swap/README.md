# v1.body_swap

## Module Functions

### Body Swap <a name="create"></a>

Swap a person into a scene image using Nano Banana 2. Credits depend on `resolution` (from 100 credits at 640px upward).

**API Endpoint**: `POST /v1/body-swap`

#### Parameters

| Parameter           | Required | Description                                                                                                                                                                                                                                                                                                                                                                | Example                                                                                                          |
| ------------------- | :------: | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `Assets`            |    ✓     | Person image and scene image for body swap                                                                                                                                                                                                                                                                                                                                 | `V1BodySwapCreateBodyAssets {PersonFilePath: "api-assets/id/1234.png",SceneFilePath: "api-assets/id/5678.png",}` |
| `└─ PersonFilePath` |    ✓     | Image of the person to place into the scene. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls). See the [file upload guide](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) for details. | `"api-assets/id/1234.png"`                                                                                       |
| `└─ SceneFilePath`  |    ✓     | Original scene image (background). This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls). See the [file upload guide](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) for details.           | `"api-assets/id/5678.png"`                                                                                       |
| `Resolution`        |    ✓     | Output resolution. Determines credits charged for the run.                                                                                                                                                                                                                                                                                                                 | `V1BodySwapCreateBodyResolutionEnum1k`                                                                           |
| `Name`              |    ✗     | Give your image a custom name for easy identification.                                                                                                                                                                                                                                                                                                                     | `"My Body Swap image"`                                                                                           |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	body_swap "github.com/magichourhq/magic-hour-go/resources/v1/body_swap"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.BodySwap.Create(body_swap.CreateRequest{
		Assets: types.V1BodySwapCreateBodyAssets{
			PersonFilePath: "api-assets/id/1234.png",
			SceneFilePath:  "api-assets/id/5678.png",
		},
		Name:       nullable.NewValue("My Body Swap image"),
		Resolution: types.V1BodySwapCreateBodyResolutionEnum1k,
	})
}
```

#### Response

##### Type

[V1BodySwapCreateResponse](/types/v1_body_swap_create_response.go)

##### Example

```go
V1BodySwapCreateResponse {
CreditsCharged: 100,
FrameCost: 100,
Id: "cuid-example",
}
```
