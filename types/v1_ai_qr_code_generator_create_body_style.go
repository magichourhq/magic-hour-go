package types

// V1AiQrCodeGeneratorCreateBodyStyle
type V1AiQrCodeGeneratorCreateBodyStyle struct {
	// To use our templates, pass in one of Watercolor, Cyberpunk City, Ink Landscape, Interior Painting, Japanese Street, Mech, Minecraft, Picasso Painting, Game Map, Spaceship, Chinese Painting, Winter Village, or pass any custom art style.
	ArtStyle string `json:"art_style"`
}
