package types

// Determines the orientation of the output video
type V1TextToVideoCreateBodyOrientationEnum string

const (
	V1TextToVideoCreateBodyOrientationEnumLandscape V1TextToVideoCreateBodyOrientationEnum = "landscape"
	V1TextToVideoCreateBodyOrientationEnumPortrait  V1TextToVideoCreateBodyOrientationEnum = "portrait"
	V1TextToVideoCreateBodyOrientationEnumSquare    V1TextToVideoCreateBodyOrientationEnum = "square"
)
