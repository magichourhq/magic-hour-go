package types

// The status of the audio.
type V1AudioProjectsGetResponseStatusEnum string

const (
	V1AudioProjectsGetResponseStatusEnumCanceled  V1AudioProjectsGetResponseStatusEnum = "canceled"
	V1AudioProjectsGetResponseStatusEnumComplete  V1AudioProjectsGetResponseStatusEnum = "complete"
	V1AudioProjectsGetResponseStatusEnumDraft     V1AudioProjectsGetResponseStatusEnum = "draft"
	V1AudioProjectsGetResponseStatusEnumError     V1AudioProjectsGetResponseStatusEnum = "error"
	V1AudioProjectsGetResponseStatusEnumQueued    V1AudioProjectsGetResponseStatusEnum = "queued"
	V1AudioProjectsGetResponseStatusEnumRendering V1AudioProjectsGetResponseStatusEnum = "rendering"
)
