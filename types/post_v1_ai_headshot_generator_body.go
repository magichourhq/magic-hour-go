// Code generated by Sideko (sideko.dev) DO NOT EDIT.

package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

type PostV1AiHeadshotGeneratorBody struct {
	Assets PostV1AiHeadshotGeneratorBodyAssets `json:"assets"`
	Name   nullable.Nullable[string]           `json:"name"`
}
