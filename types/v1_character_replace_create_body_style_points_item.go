package types

// V1CharacterReplaceCreateBodyStylePointsItem
type V1CharacterReplaceCreateBodyStylePointsItem struct {
	// Horizontal pixel coordinate in the source video frame at `time_seconds`, measured from the left edge.
	PositionX int `json:"position_x"`
	// Vertical pixel coordinate in the source video frame at `time_seconds`, measured from the top edge.
	PositionY int `json:"position_y"`
	// Timestamp on the source video timeline in seconds. Uses the same clock as `start_seconds` and `end_seconds`.
	TimeSeconds float64 `json:"time_seconds"`
}
