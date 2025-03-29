package types

// The status of the video.
type V1VideoProjectsGetResponseStatusEnum string

const (
	V1VideoProjectsGetResponseStatusEnumCanceled  V1VideoProjectsGetResponseStatusEnum = "canceled"
	V1VideoProjectsGetResponseStatusEnumComplete  V1VideoProjectsGetResponseStatusEnum = "complete"
	V1VideoProjectsGetResponseStatusEnumDraft     V1VideoProjectsGetResponseStatusEnum = "draft"
	V1VideoProjectsGetResponseStatusEnumError     V1VideoProjectsGetResponseStatusEnum = "error"
	V1VideoProjectsGetResponseStatusEnumQueued    V1VideoProjectsGetResponseStatusEnum = "queued"
	V1VideoProjectsGetResponseStatusEnumRendering V1VideoProjectsGetResponseStatusEnum = "rendering"
)
