package types

// Face editing parameters
type V1AiFaceEditorCreateBodyStyle struct {
	// Enhance face features
	EnhanceFace bool `json:"enhance_face"`
	// Horizontal eye gaze (-100 to 100), in increments of 5
	EyeGazeHorizontal float64 `json:"eye_gaze_horizontal"`
	// Vertical eye gaze (-100 to 100), in increments of 5
	EyeGazeVertical float64 `json:"eye_gaze_vertical"`
	// Eye open ratio (-100 to 100), in increments of 5
	EyeOpenRatio float64 `json:"eye_open_ratio"`
	// Eyebrow direction (-100 to 100), in increments of 5
	EyebrowDirection float64 `json:"eyebrow_direction"`
	// Head pitch (-100 to 100), in increments of 5
	HeadPitch float64 `json:"head_pitch"`
	// Head roll (-100 to 100), in increments of 5
	HeadRoll float64 `json:"head_roll"`
	// Head yaw (-100 to 100), in increments of 5
	HeadYaw float64 `json:"head_yaw"`
	// Lip open ratio (-100 to 100), in increments of 5
	LipOpenRatio float64 `json:"lip_open_ratio"`
	// Mouth grim (-100 to 100), in increments of 5
	MouthGrim float64 `json:"mouth_grim"`
	// Horizontal mouth position (-100 to 100), in increments of 5
	MouthPositionHorizontal float64 `json:"mouth_position_horizontal"`
	// Vertical mouth position (-100 to 100), in increments of 5
	MouthPositionVertical float64 `json:"mouth_position_vertical"`
	// Mouth pout (-100 to 100), in increments of 5
	MouthPout float64 `json:"mouth_pout"`
	// Mouth purse (-100 to 100), in increments of 5
	MouthPurse float64 `json:"mouth_purse"`
	// Mouth smile (-100 to 100), in increments of 5
	MouthSmile float64 `json:"mouth_smile"`
}
