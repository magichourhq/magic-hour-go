# v1.estimate.image

## Module Functions

### Estimate image credits <a name="list"></a>

Get the credits an image create request charges. Pass `tool`, then the fields below that the tool's create request uses, with the same values. Does not create a project or spend credits. It prices the fields passed here. It does not check your plan, your balance or the rest of the create request.

**API Endpoint**: `GET /v1/estimate/image`

#### Parameters

| Parameter     | Required | Description                                                                                                                             | Example                                    |
| ------------- | :------: | --------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------ |
| `Tool`        |    ✓     | The create endpoint to price, named by its path. For example `ai-image-editor` for `POST /v1/ai-image-editor`.                          | `V1EstimateImageListToolEnumAiImageEditor` |
| `ImageCount`  |    ✗     | The `image_count` of the create request. Used by `ai-image-generator`, `ai-image-editor`. Defaults to 1.                                | `4`                                        |
| `Mode`        |    ✗     | The `style.mode` of the create request. Used by `ai-image-upscaler`. One of `preserve`, `balanced`, `creative`. Defaults to `balanced`. | `"string"`                                 |
| `Model`       |    ✗     | The `model` of the create request. Used by `ai-image-generator`, `ai-image-editor`.                                                     | `"nano-banana-2.1"`                        |
| `Resolution`  |    ✗     | The `resolution` of the create request. Used by `ai-image-generator`, `ai-image-editor`, `body-swap`. One of `640px`, `1k`, `2k`, `4k`. | `"2k"`                                     |
| `ScaleFactor` |    ✗     | The `scale_factor` of the create request. Used by `ai-image-upscaler`. One of `2`, `4`.                                                 | `123.0`                                    |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	image "github.com/magichourhq/magic-hour-go/resources/v1/estimate/image"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.Estimate.Image.List(image.ListRequest{
		Tool: types.V1EstimateImageListToolEnumAiImageEditor,
	})
}
```

#### Response

##### Type

[V1EstimateImageListResponse](/types/v1_estimate_image_list_response.go)

##### Example

```go
V1EstimateImageListResponse {
EstimatedCredits: 150,
}
```
