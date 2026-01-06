# v1.video_projects

## Module Functions

### Delete video <a name="delete"></a>

Permanently delete the rendered video. This action is not reversible, please be sure before deleting.

**API Endpoint**: `DELETE /v1/video-projects/{id}`

#### Parameters

| Parameter | Required | Description                                                                                         | Example          |
| --------- | :------: | --------------------------------------------------------------------------------------------------- | ---------------- |
| `Id`      |    ✓     | Unique ID of the video project. This value is returned by all of the POST APIs that create a video. | `"cuid-example"` |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	video_projects "github.com/magichourhq/magic-hour-go/resources/v1/video_projects"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	err := client.V1.VideoProjects.Delete(video_projects.DeleteRequest{
		Id: "cuid-example",
	})
}
```

### Get video details <a name="get"></a>

Check the progress of a video project. The `downloads` field is populated after a successful render.

**Statuses**

- `queued` — waiting to start
- `rendering` — in progress
- `complete` — ready; see `downloads`
- `error` — a failure occurred (see `error`)
- `canceled` — user canceled
- `draft` — not used

**API Endpoint**: `GET /v1/video-projects/{id}`

#### Parameters

| Parameter | Required | Description                                                                                         | Example          |
| --------- | :------: | --------------------------------------------------------------------------------------------------- | ---------------- |
| `Id`      |    ✓     | Unique ID of the video project. This value is returned by all of the POST APIs that create a video. | `"cuid-example"` |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	video_projects "github.com/magichourhq/magic-hour-go/resources/v1/video_projects"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.VideoProjects.Get(video_projects.GetRequest{
		Id: "cuid-example",
	})
}
```

#### Response

##### Type

[V1VideoProjectsGetResponse](/types/v1_video_projects_get_response.go)

##### Example

```go
V1VideoProjectsGetResponse {
CreatedAt: "1970-01-01T00:00:00",
CreditsCharged: 450,
Download: nullable.NewValue(V1VideoProjectsGetResponseDownload {
ExpiresAt: "2024-10-19T05:16:19.027Z",
Url: "https://videos.magichour.ai/id/output.mp4",
}),
Downloads: []V1VideoProjectsGetResponseDownloadsItem{
V1VideoProjectsGetResponseDownloadsItem {
ExpiresAt: "2024-10-19T05:16:19.027Z",
Url: "https://videos.magichour.ai/id/output.mp4",
},
},
Enabled: true,
EndSeconds: 15.0,
Error: nullable.NewValue(V1VideoProjectsGetResponseError {
Code: "no_source_face",
Message: "Please use an image with a detectable face",
}),
Fps: 30.0,
Height: 960,
Id: "cuid-example",
Name: nullable.NewValue("Example Name"),
StartSeconds: 0.0,
Status: V1VideoProjectsGetResponseStatusEnumComplete,
TotalFrameCost: 450,
Type: "FACE_SWAP",
Width: 512,
}
```
