package types

// * `Dreamshaper` - a good all-around model that works for both animations as well as realism.
// * `Absolute Reality` - better at realism, but you'll often get similar results with Dreamshaper as well.
// * `Flat 2D Anime` - best for a flat illustration style that's common in most anime.
// * `default` - use the default recommended model for the selected art style.
type V1VideoToVideoCreateBodyStyleModelEnum string

const (
	V1VideoToVideoCreateBodyStyleModelEnum3dAnime         V1VideoToVideoCreateBodyStyleModelEnum = "3D Anime"
	V1VideoToVideoCreateBodyStyleModelEnumAbsoluteReality V1VideoToVideoCreateBodyStyleModelEnum = "Absolute Reality"
	V1VideoToVideoCreateBodyStyleModelEnumDreamshaper     V1VideoToVideoCreateBodyStyleModelEnum = "Dreamshaper"
	V1VideoToVideoCreateBodyStyleModelEnumFlat2dAnime     V1VideoToVideoCreateBodyStyleModelEnum = "Flat 2D Anime"
	V1VideoToVideoCreateBodyStyleModelEnumKaywaii         V1VideoToVideoCreateBodyStyleModelEnum = "Kaywaii"
	V1VideoToVideoCreateBodyStyleModelEnumSoftAnime       V1VideoToVideoCreateBodyStyleModelEnum = "Soft Anime"
	V1VideoToVideoCreateBodyStyleModelEnumWesternAnime    V1VideoToVideoCreateBodyStyleModelEnum = "Western Anime"
	V1VideoToVideoCreateBodyStyleModelEnumDefault         V1VideoToVideoCreateBodyStyleModelEnum = "default"
)
