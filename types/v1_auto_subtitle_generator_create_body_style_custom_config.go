package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// Custom subtitle configuration.
type V1AutoSubtitleGeneratorCreateBodyStyleCustomConfig struct {
	// Font name from Google Fonts. Not all fonts support all languages or character sets.
	// We recommend verifying language support and appearance directly on https://fonts.google.com before use.
	Font nullable.Nullable[string] `json:"font,omitempty"`
	// Font size in pixels. If not provided, the font size is automatically calculated based on the video resolution.
	FontSize nullable.Nullable[float64] `json:"font_size,omitempty"`
	// Font style (e.g., normal, italic, bold)
	FontStyle nullable.Nullable[string] `json:"font_style,omitempty"`
	// Color used to highlight the current spoken text
	HighlightedTextColor nullable.Nullable[string] `json:"highlighted_text_color,omitempty"`
	// Horizontal alignment of the text (e.g., left, center, right)
	HorizontalPosition nullable.Nullable[string] `json:"horizontal_position,omitempty"`
	// Stroke (outline) color of the text
	StrokeColor nullable.Nullable[string] `json:"stroke_color,omitempty"`
	// Width of the text stroke in pixels. If `stroke_color` is provided, but `stroke_width` is not, the `stroke_width` will be calculated automatically based on the font size.
	StrokeWidth nullable.Nullable[float64] `json:"stroke_width,omitempty"`
	// Primary text color in hex format
	TextColor nullable.Nullable[string] `json:"text_color,omitempty"`
	// Vertical alignment of the text (e.g., top, center, bottom)
	VerticalPosition nullable.Nullable[string] `json:"vertical_position,omitempty"`
}
