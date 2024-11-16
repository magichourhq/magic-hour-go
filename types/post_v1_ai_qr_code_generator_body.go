// Code generated by Sideko (sideko.dev) DO NOT EDIT.

package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

type PostV1AiQrCodeGeneratorBody struct {
	Content string                           `json:"content"`
	Name    nullable.Nullable[string]        `json:"name,omitempty"`
	Style   PostV1AiQrCodeGeneratorBodyStyle `json:"style"`
}
