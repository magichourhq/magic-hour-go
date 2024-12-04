
# Magic Hour API Go SDK

## Overview

# Introduction 

Magic Hour provides an API (beta) that can be integrated into your own application to generate videos using AI. 

Webhook documentation can be found [here](https://magichour.ai/docs/webhook).

If you have any questions, please reach out to us via [discord](https://discord.gg/JX5rgsZaJp).

# Authentication

Every request requires an API key.

To get started, first generate your API key [here](https://magichour.ai/settings/developer).

Then, add the `Authorization` header to the request.

| Key | Value |
|-|-|
| Authorization | Bearer mhk_live_apikey |

> **Warning**: any API call that renders a video will utilize frames in your account.


### Example Client Initialization

```go
import (
	sdk "github.com/magichourhq/magic-hour-go/client"
	os "os"
)

client := sdk.NewClient(sdk.WithBearerAuth(os.Getenv("API_TOKEN")))
```

## Module Documentation and Snippets

### [v1.ai_headshot_generator](resources/v1/ai_headshot_generator/README.md)

* [create](resources/v1/ai_headshot_generator/README.md#create) - Create AI Headshots

### [v1.ai_image_generator](resources/v1/ai_image_generator/README.md)

* [create](resources/v1/ai_image_generator/README.md#create) - Create AI Images

### [v1.ai_image_upscaler](resources/v1/ai_image_upscaler/README.md)

* [create](resources/v1/ai_image_upscaler/README.md#create) - Create Upscaled Image

### [v1.ai_photo_editor](resources/v1/ai_photo_editor/README.md)

* [create](resources/v1/ai_photo_editor/README.md#create) - AI Photo Editor

### [v1.ai_qr_code_generator](resources/v1/ai_qr_code_generator/README.md)

* [create](resources/v1/ai_qr_code_generator/README.md#create) - Create AI QR Code

### [v1.face_swap](resources/v1/face_swap/README.md)

* [create](resources/v1/face_swap/README.md#create) - Create Face Swap video

### [v1.face_swap_photo](resources/v1/face_swap_photo/README.md)

* [create](resources/v1/face_swap_photo/README.md#create) - Create Face Swap Photo

### [v1.files.upload_urls](resources/v1/files/upload_urls/README.md)

* [create](resources/v1/files/upload_urls/README.md#create) - Generate asset upload urls

### [v1.image_projects](resources/v1/image_projects/README.md)

* [delete](resources/v1/image_projects/README.md#delete) - Delete image
* [get](resources/v1/image_projects/README.md#get) - Get image project details

### [v1.image_to_video](resources/v1/image_to_video/README.md)

* [create](resources/v1/image_to_video/README.md#create) - Create Image-to-Video

### [v1.lip_sync](resources/v1/lip_sync/README.md)

* [create](resources/v1/lip_sync/README.md#create) - Create Lip Sync video

### [v1.text_to_video](resources/v1/text_to_video/README.md)

* [create](resources/v1/text_to_video/README.md#create) - Create Text-to-Video

### [v1.video_projects](resources/v1/video_projects/README.md)

* [delete](resources/v1/video_projects/README.md#delete) - Delete video
* [get](resources/v1/video_projects/README.md#get) - Get video details

### [v1.video_to_video](resources/v1/video_to_video/README.md)

* [create](resources/v1/video_to_video/README.md#create) - Create Video-to-Video

<!-- MODULE DOCS END -->
