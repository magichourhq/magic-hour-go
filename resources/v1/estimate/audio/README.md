# v1.estimate.audio

## Module Functions

### Estimate audio credits <a name="list"></a>

Get the credits an audio create request charges. Pass `tool`, then the field below that the tool's create request uses. Does not create a project or spend credits. It prices the fields passed here. It does not check your plan, your balance or the rest of the create request.

**API Endpoint**: `GET /v1/estimate/audio`

#### Parameters

| Parameter        | Required | Description                                                                                                            | Example                                       |
| ---------------- | :------: | ---------------------------------------------------------------------------------------------------------------------- | --------------------------------------------- |
| `Tool`           |    ✓     | The create endpoint to price, named by its path. For example `ai-voice-generator` for `POST /v1/ai-voice-generator`.   | `V1EstimateAudioListToolEnumAiVoiceGenerator` |
| `CharacterCount` |    ✗     | The number of characters in the `style.prompt` of the create request. Used by `ai-voice-generator`, `ai-voice-cloner`. | `120`                                         |
| `Duration`       |    ✗     | The `duration` of the create request. Used by `ai-music-generator`.                                                    | `123`                                         |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	audio "github.com/magichourhq/magic-hour-go/resources/v1/estimate/audio"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.Estimate.Audio.List(audio.ListRequest{
		Tool: types.V1EstimateAudioListToolEnumAiVoiceGenerator,
	})
}
```

#### Response

##### Type

[V1EstimateAudioListResponse](/types/v1_estimate_audio_list_response.go)

##### Example

```go
V1EstimateAudioListResponse {
EstimatedCredits: 150,
}
```
