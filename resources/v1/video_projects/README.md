
### delete <a name="delete"></a>
Delete video

Permanently delete the rendered video. This action is not reversible, please be sure before deleting.

**API Endpoint**: `DELETE /v1/video-projects/{id}`

#### Example Snippet

```go
import (
	sdk "github.com/magichourhq/magic-hour-go/client"
	os "os"
	video_projects "github.com/magichourhq/magic-hour-go/resources/v1/video_projects"
)

client := sdk.NewClient(sdk.WithBearerAuth(os.Getenv("API_TOKEN")))
err := client.V1.VideoProjects.Delete(video_projects.DeleteRequest { Id: "string" })
```

### get <a name="get"></a>
Get video details

Get the details of a video project. The `download` field will be `null` unless the video was successfully rendered.

The video can be one of the following status
- `draft` - not currently used
- `queued` - the job is queued and waiting for a GPU
- `rendering` - the generation is in progress
- `complete` - the video is successful created
- `error` - an error occurred during rendering
- `canceled` - video render is canceled by the user


**API Endpoint**: `GET /v1/video-projects/{id}`

#### Example Snippet

```go
import (
	sdk "github.com/magichourhq/magic-hour-go/client"
	os "os"
	video_projects "github.com/magichourhq/magic-hour-go/resources/v1/video_projects"
)

client := sdk.NewClient(sdk.WithBearerAuth(os.Getenv("API_TOKEN")))
res, err := client.V1.VideoProjects.Get(video_projects.GetRequest { Id: "string" })
```
