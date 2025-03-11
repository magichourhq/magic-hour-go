
package types



// * `Dreamshaper` - a good all-around model that works for both animations as well as realism. 
// * `Absolute Reality` - better at realism, but you'll often get similar results with Dreamshaper as well. 
// * `Flat 2D Anime` - best for a flat illustration style that's common in most anime.
// * `default` - use the default recommended model for the selected art style.
type PostV1VideoToVideoBodyStyleModelEnum string
const (
    PostV1VideoToVideoBodyStyleModelEnumAbsoluteReality PostV1VideoToVideoBodyStyleModelEnum = "Absolute Reality"
    PostV1VideoToVideoBodyStyleModelEnumDreamshaper PostV1VideoToVideoBodyStyleModelEnum = "Dreamshaper"
    PostV1VideoToVideoBodyStyleModelEnumFlat2dAnime PostV1VideoToVideoBodyStyleModelEnum = "Flat 2D Anime"
    PostV1VideoToVideoBodyStyleModelEnumDefault PostV1VideoToVideoBodyStyleModelEnum = "default"
)


