# v1.ai_music_generator

## Module Functions

### AI Music Generator <a name="create"></a>

Generate music from a description, with optional custom lyrics or instrumental-only output. Costs 1 credits per requested second. Returns a project ID; poll GET /v1/audio-projects/{id} for status and downloads.

**API Endpoint**: `POST /v1/ai-music-generator`

#### Parameters

| Parameter              | Required | Description                                                                            | Example                                                                                                                                                                                        |
| ---------------------- | :------: | -------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Duration`             |    ✓     | Requested track duration in seconds, from 15 to 300. Determines credits charged.       | `120`                                                                                                                                                                                          |
| `Style`                |    ✓     |                                                                                        | `V1AiMusicGeneratorCreateBodyStyle {Instrumental: nullable.NewValue(false),Lyrics: nullable.NewValue("City lights are fading..."),Prompt: "Dreamy synthwave with warm bass and soft vocals",}` |
| `└─ Instrumental`      |    ✗     | Generate music without vocals. When true, lyrics is ignored.                           | `false`                                                                                                                                                                                        |
| `└─ Lyrics`            |    ✗     | Custom lyrics. Omit or leave blank for AI-generated lyrics when instrumental is false. | `"City lights are fading..."`                                                                                                                                                                  |
| `└─ Prompt`            |    ✓     | Describe the song, including genre, mood, and instruments.                             | `"Dreamy synthwave with warm bass and soft vocals"`                                                                                                                                            |
| `AdditionalProperties` |    ✗     |                                                                                        | `make(map[string]interface{})`                                                                                                                                                                 |
| `Name`                 |    ✗     | Give your audio a custom name for easy identification.                                 | `"My Music Generator audio"`                                                                                                                                                                   |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	ai_music_generator "github.com/magichourhq/magic-hour-go/resources/v1/ai_music_generator"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.AiMusicGenerator.Create(ai_music_generator.CreateRequest{
		Duration: 120,
		Name:     nullable.NewValue("My Music Generator audio"),
		Style: types.V1AiMusicGeneratorCreateBodyStyle{
			Instrumental: nullable.NewValue(false),
			Lyrics:       nullable.NewValue("City lights are fading..."),
			Prompt:       "Dreamy synthwave with warm bass and soft vocals",
		},
	})
}
```

#### Response

##### Type

[V1AiMusicGeneratorCreateResponse](/types/v1_ai_music_generator_create_response.go)

##### Example

```go
V1AiMusicGeneratorCreateResponse {
CreditsCharged: 120,
Id: "cuid-example",
}
```
