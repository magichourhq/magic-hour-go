package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// Face editing parameters
type V1AiFaceEditorCreateBodyStyle struct {
	// Enhance face features
	EnhanceFace nullable.Nullable[bool] `json:"enhance_face,omitempty"`
	// Horizontal eye gaze (-100 to 100), in increments of 5
	EyeGazeHorizontal nullable.Nullable[float64] `json:"eye_gaze_horizontal,omitempty"`
	// Vertical eye gaze (-100 to 100), in increments of 5
	EyeGazeVertical nullable.Nullable[float64] `json:"eye_gaze_vertical,omitempty"`
	// Eye open ratio (-100 to 100), in increments of 5
	EyeOpenRatio nullable.Nullable[float64] `json:"eye_open_ratio,omitempty"`
	// Eyebrow direction (-100 to 100), in increments of 5
	EyebrowDirection nullable.Nullable[float64] `json:"eyebrow_direction,omitempty"`
	// Head pitch (-100 to 100), in increments of 5
	HeadPitch nullable.Nullable[float64] `json:"head_pitch,omitempty"`
	// Head roll (-100 to 100), in increments of 5
	HeadRoll nullable.Nullable[float64] `json:"head_roll,omitempty"`
	// Head yaw (-100 to 100), in increments of 5
	HeadYaw nullable.Nullable[float64] `json:"head_yaw,omitempty"`
	// Lip open ratio (-100 to 100), in increments of 5
	LipOpenRatio nullable.Nullable[float64] `json:"lip_open_ratio,omitempty"`
	// Mouth grim (-100 to 100), in increments of 5
	MouthGrim nullable.Nullable[float64] `json:"mouth_grim,omitempty"`
	// Horizontal mouth position (-100 to 100), in increments of 5
	MouthPositionHorizontal nullable.Nullable[float64] `json:"mouth_position_horizontal,omitempty"`
	// Vertical mouth position (-100 to 100), in increments of 5
	MouthPositionVertical nullable.Nullable[float64] `json:"mouth_position_vertical,omitempty"`
	// Mouth pout (-100 to 100), in increments of 5
	MouthPout nullable.Nullable[float64] `json:"mouth_pout,omitempty"`
	// Mouth purse (-100 to 100), in increments of 5
	MouthPurse nullable.Nullable[float64] `json:"mouth_purse,omitempty"`
	// Mouth smile (-100 to 100), in increments of 5
	MouthSmile nullable.Nullable[float64] `json:"mouth_smile,omitempty"`
}
