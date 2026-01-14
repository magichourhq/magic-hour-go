package types

// Type of garment to swap. If not provided, swaps the entire outfit.
// * `upper_body` - for shirts/jackets
// * `lower_body` - for pants/skirts
// * `dresses` - for entire outfit (deprecated, use `entire_outfit` instead)
// * `entire_outfit` - for entire outfit
type V1AiClothesChangerCreateBodyAssetsGarmentTypeEnum string

const (
	V1AiClothesChangerCreateBodyAssetsGarmentTypeEnumDresses      V1AiClothesChangerCreateBodyAssetsGarmentTypeEnum = "dresses"
	V1AiClothesChangerCreateBodyAssetsGarmentTypeEnumEntireOutfit V1AiClothesChangerCreateBodyAssetsGarmentTypeEnum = "entire_outfit"
	V1AiClothesChangerCreateBodyAssetsGarmentTypeEnumLowerBody    V1AiClothesChangerCreateBodyAssetsGarmentTypeEnum = "lower_body"
	V1AiClothesChangerCreateBodyAssetsGarmentTypeEnumUpperBody    V1AiClothesChangerCreateBodyAssetsGarmentTypeEnum = "upper_body"
)
