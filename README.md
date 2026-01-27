# Magic Hour Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/magichourhq/magic-hour-go.svg)](https://pkg.go.dev/github.com/magichourhq/magic-hour-go)

Magic Hour provides an API (beta) that can be integrated into your own application to generate videos and images using AI.

Webhook documentation can be found [here](https://magichour.ai/docs/webhook).

If you have any questions, please reach out to us via [discord](https://discord.gg/JX5rgsZaJp).

## Install

```sh
go get -u github.com/magichourhq/magic-hour-go
```

## Usage

Initialize the client

```go
import (
	sdk "github.com/magichourhq/magic-hour-go/client"
)

// generate your API Key at https://magichour.ai/developer
client := sdk.NewClient(sdk.WithBearerAuth("my api key"))
```

> [!WARNING]
> Any API call that renders a video will utilize frames in your account.

## Module Documentation and Snippets

### [V1.AiClothesChanger](resources/v1/ai_clothes_changer/README.md)

- [create](resources/v1/ai_clothes_changer/README.md#create) - AI Clothes Changer

### [V1.AiFaceEditor](resources/v1/ai_face_editor/README.md)

- [create](resources/v1/ai_face_editor/README.md#create) - AI Face Editor

### [V1.AiGifGenerator](resources/v1/ai_gif_generator/README.md)

- [create](resources/v1/ai_gif_generator/README.md#create) - AI GIF Generator

### [V1.AiHeadshotGenerator](resources/v1/ai_headshot_generator/README.md)

- [create](resources/v1/ai_headshot_generator/README.md#create) - AI Headshot Generator

### [V1.AiImageEditor](resources/v1/ai_image_editor/README.md)

- [create](resources/v1/ai_image_editor/README.md#create) - AI Image Editor

### [V1.AiImageGenerator](resources/v1/ai_image_generator/README.md)

- [create](resources/v1/ai_image_generator/README.md#create) - AI Image Generator

### [V1.AiImageUpscaler](resources/v1/ai_image_upscaler/README.md)

- [create](resources/v1/ai_image_upscaler/README.md#create) - AI Image Upscaler

### [V1.AiMemeGenerator](resources/v1/ai_meme_generator/README.md)

- [create](resources/v1/ai_meme_generator/README.md#create) - AI Meme Generator

### [V1.AiQrCodeGenerator](resources/v1/ai_qr_code_generator/README.md)

- [create](resources/v1/ai_qr_code_generator/README.md#create) - AI QR Code Generator

### [V1.AiTalkingPhoto](resources/v1/ai_talking_photo/README.md)

- [create](resources/v1/ai_talking_photo/README.md#create) - AI Talking Photo

### [V1.AiVoiceCloner](resources/v1/ai_voice_cloner/README.md)

- [create](resources/v1/ai_voice_cloner/README.md#create) - AI Voice Cloner

### [V1.AiVoiceGenerator](resources/v1/ai_voice_generator/README.md)

- [create](resources/v1/ai_voice_generator/README.md#create) - AI Voice Generator

### [V1.Animation](resources/v1/animation/README.md)

- [create](resources/v1/animation/README.md#create) - Animation

### [V1.AudioProjects](resources/v1/audio_projects/README.md)

- [delete](resources/v1/audio_projects/README.md#delete) - Delete audio
- [get](resources/v1/audio_projects/README.md#get) - Get audio details

### [V1.AutoSubtitleGenerator](resources/v1/auto_subtitle_generator/README.md)

- [create](resources/v1/auto_subtitle_generator/README.md#create) - Auto Subtitle Generator

### [V1.FaceDetection](resources/v1/face_detection/README.md)

- [create](resources/v1/face_detection/README.md#create) - Face Detection
- [get](resources/v1/face_detection/README.md#get) - Get face detection details

### [V1.FaceSwap](resources/v1/face_swap/README.md)

- [create](resources/v1/face_swap/README.md#create) - Face Swap Video

### [V1.FaceSwapPhoto](resources/v1/face_swap_photo/README.md)

- [create](resources/v1/face_swap_photo/README.md#create) - Face Swap Photo

### [V1.Files.UploadUrls](resources/v1/files/upload_urls/README.md)

- [create](resources/v1/files/upload_urls/README.md#create) - Generate asset upload urls

### [V1.ImageBackgroundRemover](resources/v1/image_background_remover/README.md)

- [create](resources/v1/image_background_remover/README.md#create) - Image Background Remover

### [V1.ImageProjects](resources/v1/image_projects/README.md)

- [delete](resources/v1/image_projects/README.md#delete) - Delete image
- [get](resources/v1/image_projects/README.md#get) - Get image details

### [V1.ImageToVideo](resources/v1/image_to_video/README.md)

- [create](resources/v1/image_to_video/README.md#create) - Image-to-Video

### [V1.LipSync](resources/v1/lip_sync/README.md)

- [create](resources/v1/lip_sync/README.md#create) - Lip Sync

### [V1.PhotoColorizer](resources/v1/photo_colorizer/README.md)

- [create](resources/v1/photo_colorizer/README.md#create) - Photo Colorizer

### [V1.TextToVideo](resources/v1/text_to_video/README.md)

- [create](resources/v1/text_to_video/README.md#create) - Text-to-Video

### [V1.VideoProjects](resources/v1/video_projects/README.md)

- [delete](resources/v1/video_projects/README.md#delete) - Delete video
- [get](resources/v1/video_projects/README.md#get) - Get video details

### [V1.VideoToVideo](resources/v1/video_to_video/README.md)

- [create](resources/v1/video_to_video/README.md#create) - Video-to-Video

<!-- MODULE DOCS END -->
