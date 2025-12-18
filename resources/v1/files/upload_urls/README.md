# v1.files.upload_urls

## Module Functions

### Generate asset upload urls <a name="create"></a>

Generates a list of pre-signed upload URLs for the assets required. This API is only necessary if you want to upload to Magic Hour's storage. Refer to the [Input Files Guide](/integration/input-files) for more details.

The response array will match the order of items in the request body.

**Valid file extensions per asset type**:
- video: mp4, m4v, mov, webm
- audio: mp3, wav, aac, flac, webm
- image: png, jpg, jpeg, heic, webp, avif, jp2, tiff, bmp
- gif: gif, webp, webm

> Note: `gif` is only supported for face swap API `video_file_path` field.

Once you receive an upload URL, send a `PUT` request to upload the file directly.

Example:

```
curl -X PUT --data '@/path/to/file/video.mp4' \
  https://videos.magichour.ai/api-assets/id/video.mp4?<auth params from the API response>
```

**API Endpoint**: `POST /v1/files/upload-urls`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `Items` | ✓ | The list of assets to upload. The response array will match the order of items in the request body. | `[]V1FilesUploadUrlsCreateBodyItemsItem{V1FilesUploadUrlsCreateBodyItemsItem {Extension: "mp4",Type: V1FilesUploadUrlsCreateBodyItemsItemTypeEnumVideo,},V1FilesUploadUrlsCreateBodyItemsItem {Extension: "mp3",Type: V1FilesUploadUrlsCreateBodyItemsItemTypeEnumAudio,},}` |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	upload_urls "github.com/magichourhq/magic-hour-go/resources/v1/files/upload_urls"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.Files.UploadUrls.Create(upload_urls.CreateRequest{
		Items: []types.V1FilesUploadUrlsCreateBodyItemsItem{
			types.V1FilesUploadUrlsCreateBodyItemsItem{
				Extension: "mp4",
				Type:      types.V1FilesUploadUrlsCreateBodyItemsItemTypeEnumVideo,
			},
			types.V1FilesUploadUrlsCreateBodyItemsItem{
				Extension: "mp3",
				Type:      types.V1FilesUploadUrlsCreateBodyItemsItemTypeEnumAudio,
			},
		},
	})
}

```

#### Response

##### Type
[V1FilesUploadUrlsCreateResponse](/types/v1_files_upload_urls_create_response.go)

##### Example
`V1FilesUploadUrlsCreateResponse {
Items: []V1FilesUploadUrlsCreateResponseItemsItem{
V1FilesUploadUrlsCreateResponseItemsItem {
ExpiresAt: "2024-07-25T16:56:21.932Z",
FilePath: "api-assets/id/video.mp4",
UploadUrl: "https://videos.magichour.ai/api-assets/id/video.mp4?auth-value=1234567890",
},
V1FilesUploadUrlsCreateResponseItemsItem {
ExpiresAt: "2024-07-25T16:56:21.932Z",
FilePath: "api-assets/id/audio.mp3",
UploadUrl: "https://videos.magichour.ai/api-assets/id/audio.mp3?auth-value=1234567890",
},
},
}`

