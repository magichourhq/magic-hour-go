# v1.voices

## Module Functions

### List voices <a name="list"></a>

List the voices available for text-to-speech. Use a voice's `id` as `style.voice_name` in the AI Voice Generator API. To fetch the next page, pass the response's next_cursor as cursor.

**API Endpoint**: `GET /v1/voices`

#### Parameters

| Parameter | Required | Description                                                        | Example    |
| --------- | :------: | ------------------------------------------------------------------ | ---------- |
| `Cursor`  |    ✗     | Opaque pagination cursor from the previous response's next_cursor. | `"string"` |
| `Limit`   |    ✗     | Maximum number of voices to return. Defaults to 20.                | `20`       |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	voices "github.com/magichourhq/magic-hour-go/resources/v1/voices"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.Voices.List(voices.ListRequest{
		Limit: nullable.NewValue(20),
	})
}
```

#### Response

##### Type

[V1VoicesListResponse](/types/v1_voices_list_response.go)

##### Example

```go
V1VoicesListResponse {
NextCursor: nullable.NewValue("string"),
Voices: []V1VoicesListResponseVoicesItem{
V1VoicesListResponseVoicesItem {
AvatarUrl: nullable.NewValue("https://d28dkohlqf5vwj.cloudfront.net/models/nano-banana/female-portrait-turtleneck.webp"),
Id: "cv-0001",
Name: "Ashley",
PreviewUrl: "https://videos.magichour.ai/presets/voices/v2/cv-0001-ashley.mp3",
},
},
}
```
