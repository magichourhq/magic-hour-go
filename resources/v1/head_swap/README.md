# v1.head_swap

## Module Functions

### Head Swap <a name="create"></a>

Swap a head onto a body image. Each image costs 10 credits. Output resolution depends on your subscription; you may set `max_resolution` lower than your plan maximum if desired.

**API Endpoint**: `POST /v1/head-swap`

#### Parameters

| Parameter         | Required | Description                                                                                                                                                                                                                                                                                                                                                           | Example                                                                                                       |
| ----------------- | :------: | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- |
| `Assets`          |    ✓     | Provide the body and head images for head swap                                                                                                                                                                                                                                                                                                                        | `V1HeadSwapCreateBodyAssets {BodyFilePath: "api-assets/id/1234.png",HeadFilePath: "api-assets/id/5678.png",}` |
| `└─ BodyFilePath` |    ✓     | Image that receives the swapped head. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls). See the [file upload guide](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) for details.   | `"api-assets/id/1234.png"`                                                                                    |
| `└─ HeadFilePath` |    ✓     | Image of the head to place on the body. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls). See the [file upload guide](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) for details. | `"api-assets/id/5678.png"`                                                                                    |
| `MaxResolution`   |    ✗     | Constrains the larger dimension (height or width) of the output. Omit to use the maximum allowed for your plan (capped at 2048px). Values above your plan maximum are clamped down to your plan's maximum.                                                                                                                                                            | `1024`                                                                                                        |
| `Name`            |    ✗     | Give your image a custom name for easy identification.                                                                                                                                                                                                                                                                                                                | `"My Head Swap image"`                                                                                        |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	head_swap "github.com/magichourhq/magic-hour-go/resources/v1/head_swap"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.HeadSwap.Create(head_swap.CreateRequest{
		Assets: types.V1HeadSwapCreateBodyAssets{
			BodyFilePath: "api-assets/id/1234.png",
			HeadFilePath: "api-assets/id/5678.png",
		},
		MaxResolution: nullable.NewValue(1024),
		Name:          nullable.NewValue("My Head Swap image"),
	})
}
```

#### Response

##### Type

[V1HeadSwapCreateResponse](/types/v1_head_swap_create_response.go)

##### Example

```go
V1HeadSwapCreateResponse {
CreditsCharged: 10,
FrameCost: 10,
Id: "cuid-example",
}
```
