package types

// The status of the video.
type V1VideoProjectsgetResponseStatusEnum string

const (
	V1VideoProjectsgetResponseStatusEnumCanceled  V1VideoProjectsgetResponseStatusEnum = "canceled"
	V1VideoProjectsgetResponseStatusEnumComplete  V1VideoProjectsgetResponseStatusEnum = "complete"
	V1VideoProjectsgetResponseStatusEnumDraft     V1VideoProjectsgetResponseStatusEnum = "draft"
	V1VideoProjectsgetResponseStatusEnumError     V1VideoProjectsgetResponseStatusEnum = "error"
	V1VideoProjectsgetResponseStatusEnumQueued    V1VideoProjectsgetResponseStatusEnum = "queued"
	V1VideoProjectsgetResponseStatusEnumRendering V1VideoProjectsgetResponseStatusEnum = "rendering"
)
