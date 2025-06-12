
### Generate asset upload urls <a name="create"></a>

Create a list of urls used to upload the assets needed to generate a video. Each video type has their own requirements on what assets are required. Please refer to the specific mode API for more details. The response array will be in the same order as the request body.

Below is the list of valid extensions for each asset type:

- video: mp4, m4v, mov, webm
- audio: mp3, mpeg, wav, aac, aiff, flac
- image: png, jpg, jpeg, webp, avif, jp2, tiff, bmp

Note: `.gif` is supported for face swap API `video_file_path` field.

After receiving the upload url, you can upload the file by sending a PUT request with the header `'Content-Type: application/octet-stream'`.

For example using curl

```
curl -X PUT -H 'Content-Type: application/octet-stream' \
  --data '@/path/to/file/video.mp4' \
  https://videos.magichour.ai/api-assets/id/video.mp4?auth-value=1234567890
```


**API Endpoint**: `POST /v1/files/upload-urls`

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

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `items` | ✓ |  | `[]V1FilesUploadUrlsCreateBodyItemsItem{V1FilesUploadUrlsCreateBodyItemsItem {Extension: "mp4",Type: V1FilesUploadUrlsCreateBodyItemsItemTypeEnumVideo,},V1FilesUploadUrlsCreateBodyItemsItem {Extension: "mp3",Type: V1FilesUploadUrlsCreateBodyItemsItemTypeEnumAudio,},}` |
