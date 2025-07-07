
### Delete video <a name="delete"></a>

Permanently delete the rendered video. This action is not reversible, please be sure before deleting.

**API Endpoint**: `DELETE /v1/video-projects/{id}`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `id` | ✓ | The id of the video project | `"cm6pvghix03bvyz0zwash6noj"` |

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
		Id: "cm6pvghix03bvyz0zwash6noj",
	})
}

```

### Get video details <a name="get"></a>

Get the details of a video project. The `downloads` field will be empty unless the video was successfully rendered.

The video can be one of the following status
- `draft` - not currently used
- `queued` - the job is queued and waiting for a GPU
- `rendering` - the generation is in progress
- `complete` - the video is successful created
- `error` - an error occurred during rendering
- `canceled` - video render is canceled by the user


**API Endpoint**: `GET /v1/video-projects/{id}`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `id` | ✓ | The id of the video | `"cm6pvghix03bvyz0zwash6noj"` |

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
		Id: "cm6pvghix03bvyz0zwash6noj",
	})
}

```

#### Response

##### Type
[V1VideoProjectsGetResponse](/types/v1_video_projects_get_response.go)

##### Example
`V1VideoProjectsGetResponse {
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
Id: "clx7uu86w0a5qp55yxz315r6r",
Name: nullable.NewValue("Example Name"),
StartSeconds: 0.0,
Status: V1VideoProjectsGetResponseStatusEnumComplete,
TotalFrameCost: 450,
Type: "FACE_SWAP",
Width: 512,
}`
