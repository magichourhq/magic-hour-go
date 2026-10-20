# v1.ai_ugc_ad

## Module Functions

### AI UGC Ad Generator <a name="create"></a>

Create a spoken UGC video from a prepared image and script. Returns a project ID; poll GET /v1/video-projects/{id} for status and downloads. To compose an image, generate the image separately with the Image Editor API and supply its download URL.

**API Endpoint**: `POST /v1/ai-ugc-ad`

#### Parameters

| Parameter              | Required | Description                                                                                                                                                                                                                                                                                                                                                                                                 | Example                                                                 |
| ---------------------- | :------: | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `Assets`               |    ✓     |                                                                                                                                                                                                                                                                                                                                                                                                             | `V1AiUgcAdCreateBodyAssets {ImageFilePath: "api-assets/id/image.jpg",}` |
| `└─ ImageFilePath`     |    ✓     | Prepared image to animate. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls). See the [file upload guide](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) for details.                                                    | `"api-assets/id/image.jpg"`                                             |
| `EndSeconds`           |    ✓     | Video duration in seconds. Supported values by model: * kling-3.0: 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15 seconds. * ltx-2.5: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60 seconds. * veo3.1: 4, 6, 8, 15, 22, 29, 36, 43, 50, 57, 64, 71, 78, 85, 92, 99, 106, 113, 120, 127, 134, 141 seconds. Free plans are limited to 5 seconds unless a credit pack has been purchased. | `4.0`                                                                   |
| `Script`               |    ✓     | The words the actor should speak. Character limits apply after trimming: * kling-3.0: 2500 characters. * ltx-2.5: 1804 characters. * veo3.1: 2500 characters. LTX's dialogue wrapper counts toward its model prompt limit.                                                                                                                                                                                  | `"Meet your new everyday essential."`                                   |
| `AdditionalProperties` |    ✗     |                                                                                                                                                                                                                                                                                                                                                                                                             | `make(map[string]interface{})`                                          |
| `Model`                |    ✗     | Defaults to kling-3.0 on paid plans and ltx-2.5 on free plans.                                                                                                                                                                                                                                                                                                                                              | `V1AiUgcAdCreateBodyModelEnumKling30`                                   |
| `Resolution`           |    ✗     | Defaults to 720p on paid plans and 480p on free plans. Unsupported model/resolution combinations and resolutions unavailable for your plan are rejected.                                                                                                                                                                                                                                                    | `V1AiUgcAdCreateBodyResolutionEnum1080p`                                |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	ai_ugc_ad "github.com/magichourhq/magic-hour-go/resources/v1/ai_ugc_ad"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.AiUgcAd.Create(ai_ugc_ad.CreateRequest{
		Assets: types.V1AiUgcAdCreateBodyAssets{
			ImageFilePath: "api-assets/id/image.jpg",
		},
		EndSeconds: 4.0,
		Script:     "Meet your new everyday essential.",
	})
}
```

#### Response

##### Type

[V1AiUgcAdCreateResponse](/types/v1_ai_ugc_ad_create_response.go)

##### Example

```go
V1AiUgcAdCreateResponse {
CreditsCharged: 450,
EstimatedFrameCost: 123,
Id: "cuid-example",
}
```
