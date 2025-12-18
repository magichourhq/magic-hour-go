# v1.face_detection

## Module Functions

### Get face detection details <a name="get"></a>

Get the details of a face detection task. 

Use this API to get the list of faces detected in the image or video to use in the [face swap photo](/api-reference/face-swap-photo/face-swap-photo) or [face swap video](/api-reference/face-swap/face-swap-video) API calls for multi-face swaps.

**API Endpoint**: `GET /v1/face-detection/{id}`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `Id` | ✓ | The id of the task. This value is returned by the [face detection API](/api-reference/files/face-detection#response-id). | `"uuid-example"` |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	face_detection "github.com/magichourhq/magic-hour-go/resources/v1/face_detection"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.FaceDetection.Get(face_detection.GetRequest{
		Id: "uuid-example",
	})
}

```

#### Response

##### Type
[V1FaceDetectionGetResponse](/types/v1_face_detection_get_response.go)

##### Example
`V1FaceDetectionGetResponse {
CreditsCharged: 0,
Faces: []V1FaceDetectionGetResponseFacesItem{
V1FaceDetectionGetResponseFacesItem {
Path: "api-assets/id/0-0.png",
Url: "https://videos.magichour.ai/api-assets/id/0-0.png",
},
},
Id: "uuid-example",
Status: V1FaceDetectionGetResponseStatusEnumComplete,
}`

### Face Detection <a name="create"></a>

Detect faces in an image or video. 
      
Use this API to get the list of faces detected in the image or video to use in the [face swap photo](/api-reference/face-swap-photo/face-swap-photo) or [face swap video](/api-reference/face-swap/face-swap-video) API calls for multi-face swaps.

Note: Face detection is free to use for the near future. Pricing may change in the future.

**API Endpoint**: `POST /v1/face-detection`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `Assets` | ✓ | Provide the assets for face detection | `V1FaceDetectionCreateBodyAssets {TargetFilePath: "api-assets/id/1234.png",}` |
| `└─ TargetFilePath` | ✓ | This is the image or video where the face will be detected. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).  Please refer to the [Input File documentation](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) to learn more.  | `"api-assets/id/1234.png"` |
| `ConfidenceScore` | ✗ | Confidence threshold for filtering detected faces.  * Higher values (e.g., 0.9) include only faces detected with high certainty, reducing false positives.  * Lower values (e.g., 0.3) include more faces, but may increase the chance of incorrect detections. | `0.5` |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	face_detection "github.com/magichourhq/magic-hour-go/resources/v1/face_detection"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.FaceDetection.Create(face_detection.CreateRequest{
		Assets: types.V1FaceDetectionCreateBodyAssets{
			TargetFilePath: "api-assets/id/1234.png",
		},
		ConfidenceScore: nullable.NewValue(0.5),
	})
}

```

#### Response

##### Type
[V1FaceDetectionCreateResponse](/types/v1_face_detection_create_response.go)

##### Example
`V1FaceDetectionCreateResponse {
CreditsCharged: 123,
Id: "uuid-example",
}`

