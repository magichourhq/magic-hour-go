package types

// Only return saved items of this type.
type V1SavedItemsListTypeEnum string

const (
	V1SavedItemsListTypeEnumBrandKit  V1SavedItemsListTypeEnum = "brand_kit"
	V1SavedItemsListTypeEnumCharacter V1SavedItemsListTypeEnum = "character"
	V1SavedItemsListTypeEnumMoodboard V1SavedItemsListTypeEnum = "moodboard"
	V1SavedItemsListTypeEnumReference V1SavedItemsListTypeEnum = "reference"
	V1SavedItemsListTypeEnumVoice     V1SavedItemsListTypeEnum = "voice"
)
