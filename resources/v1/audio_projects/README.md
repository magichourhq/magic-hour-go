# v1.audio_projects

## Module Functions

### Delete audio <a name="delete"></a>

Permanently delete the rendered audio file(s). This action is not reversible, please be sure before deleting.

**API Endpoint**: `DELETE /v1/audio-projects/{id}`

#### Parameters

| Parameter | Required | Description                                                                                          | Example          |
| --------- | :------: | ---------------------------------------------------------------------------------------------------- | ---------------- |
| `Id`      |    ✓     | Unique ID of the audio project. This value is returned by all of the POST APIs that create an audio. | `"cuid-example"` |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	audio_projects "github.com/magichourhq/magic-hour-go/resources/v1/audio_projects"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	err := client.V1.AudioProjects.Delete(audio_projects.DeleteRequest{
		Id: "cuid-example",
	})
}
```

### Get audio details <a name="get"></a>

Check the progress of a audio project. The `downloads` field is populated after a successful render.

**Statuses**

- `queued` — waiting to start
- `rendering` — in progress
- `complete` — ready; see `downloads`
- `error` — a failure occurred (see `error`)
- `canceled` — user canceled
- `draft` — not used

**API Endpoint**: `GET /v1/audio-projects/{id}`

#### Parameters

| Parameter | Required | Description                                                                                          | Example          |
| --------- | :------: | ---------------------------------------------------------------------------------------------------- | ---------------- |
| `Id`      |    ✓     | Unique ID of the audio project. This value is returned by all of the POST APIs that create an audio. | `"cuid-example"` |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	audio_projects "github.com/magichourhq/magic-hour-go/resources/v1/audio_projects"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.AudioProjects.Get(audio_projects.GetRequest{
		Id: "cuid-example",
	})
}
```

#### Response

##### Type

[V1AudioProjectsGetResponse](/types/v1_audio_projects_get_response.go)

##### Example

```go
V1AudioProjectsGetResponse {
CreatedAt: "1970-01-01T00:00:00",
CreditsCharged: 2,
Downloads: []V1AudioProjectsGetResponseDownloadsItem{
V1AudioProjectsGetResponseDownloadsItem {
ExpiresAt: "2024-10-19T05:16:19.027Z",
Url: "https://videos.magichour.ai/id/output.wav",
},
},
Enabled: true,
Error: nullable.NewValue(V1AudioProjectsGetResponseError {
Code: "no_source_face",
Message: "Please use an image with a detectable face",
}),
Id: "cuid-example",
Name: nullable.NewValue("Example Name"),
Status: V1AudioProjectsGetResponseStatusEnumComplete,
Type: "VOICE_GENERATOR",
}
```
