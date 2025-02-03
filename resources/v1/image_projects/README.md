
### delete <a name="delete"></a>
Delete image

Permanently delete the rendered image. This action is not reversible, please be sure before deleting.

**API Endpoint**: `DELETE /v1/image-projects/{id}`

#### Example Snippet

```go
import (
	sdk "github.com/magichourhq/magic-hour-go/client"
	os "os"
	image_projects "github.com/magichourhq/magic-hour-go/resources/v1/image_projects"
)

client := sdk.NewClient(sdk.WithBearerAuth(os.Getenv("API_TOKEN")))
err := client.V1.ImageProjects.Delete(image_projects.DeleteRequest { Id: "string" })
```

### get <a name="get"></a>
Get image details

Get the details of a image project. The `download` field will be `null` unless the image was successfully rendered.

The image can be one of the following status
- `draft` - not currently used
- `queued` - the job is queued and waiting for a GPU
- `rendering` - the generation is in progress
- `complete` - the image is successful created
- `error` - an error occurred during rendering
- `canceled` - image render is canceled by the user


**API Endpoint**: `GET /v1/image-projects/{id}`

#### Example Snippet

```go
import (
	sdk "github.com/magichourhq/magic-hour-go/client"
	os "os"
	image_projects "github.com/magichourhq/magic-hour-go/resources/v1/image_projects"
)

client := sdk.NewClient(sdk.WithBearerAuth(os.Getenv("API_TOKEN")))
res, err := client.V1.ImageProjects.Get(image_projects.GetRequest { Id: "string" })
```
