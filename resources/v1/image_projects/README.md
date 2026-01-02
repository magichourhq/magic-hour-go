# v1.image_projects

## Module Functions

### Delete image <a name="delete"></a>

Permanently delete the rendered image(s). This action is not reversible, please be sure before deleting.

**API Endpoint**: `DELETE /v1/image-projects/{id}`

#### Parameters

| Parameter | Required | Description                                                                                          | Example          |
| --------- | :------: | ---------------------------------------------------------------------------------------------------- | ---------------- |
| `Id`      |    ✓     | Unique ID of the image project. This value is returned by all of the POST APIs that create an image. | `"cuid-example"` |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	image_projects "github.com/magichourhq/magic-hour-go/resources/v1/image_projects"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	err := client.V1.ImageProjects.Delete(image_projects.DeleteRequest{
		Id: "cuid-example",
	})
}
```

### Get image details <a name="get"></a>

Get the details of a image project. The `downloads` field will be empty unless the image was successfully rendered.

The image can be one of the following status

- `draft` - not currently used
- `queued` - the job is queued and waiting for a GPU
- `rendering` - the generation is in progress
- `complete` - the image is successful created
- `error` - an error occurred during rendering
- `canceled` - image render is canceled by the user

**API Endpoint**: `GET /v1/image-projects/{id}`

#### Parameters

| Parameter | Required | Description                                                                                          | Example          |
| --------- | :------: | ---------------------------------------------------------------------------------------------------- | ---------------- |
| `Id`      |    ✓     | Unique ID of the image project. This value is returned by all of the POST APIs that create an image. | `"cuid-example"` |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	image_projects "github.com/magichourhq/magic-hour-go/resources/v1/image_projects"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.ImageProjects.Get(image_projects.GetRequest{
		Id: "cuid-example",
	})
}
```

#### Response

##### Type

[V1ImageProjectsGetResponse](/types/v1_image_projects_get_response.go)

##### Example

```go
V1ImageProjectsGetResponse {
CreatedAt: "1970-01-01T00:00:00",
CreditsCharged: 5,
Downloads: []V1ImageProjectsGetResponseDownloadsItem{
V1ImageProjectsGetResponseDownloadsItem {
ExpiresAt: "2024-10-19T05:16:19.027Z",
Url: "https://videos.magichour.ai/id/output.png",
},
},
Enabled: true,
Error: nullable.NewValue(V1ImageProjectsGetResponseError {
Code: "no_source_face",
Message: "Please use an image with a detectable face",
}),
Id: "cuid-example",
ImageCount: 1,
Name: nullable.NewValue("Example Name"),
Status: V1ImageProjectsGetResponseStatusEnumComplete,
TotalFrameCost: 5,
Type: "AI_IMAGE",
}
```
