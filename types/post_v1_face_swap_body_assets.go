// Code generated by Sideko (sideko.dev) DO NOT EDIT.

package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

type PostV1FaceSwapBodyAssets struct {
	ImageFilePath string                                  `json:"image_file_path"`
	VideoFilePath nullable.Nullable[string]               `json:"video_file_path,omitempty"`
	VideoSource   PostV1FaceSwapBodyAssetsVideoSourceEnum `json:"video_source"`
	YoutubeUrl    nullable.Nullable[string]               `json:"youtube_url,omitempty"`
}
