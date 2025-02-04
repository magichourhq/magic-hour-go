package types

// The status of the image.
type V1ImageProjectsgetResponseStatusEnum string

const (
	V1ImageProjectsgetResponseStatusEnumCanceled  V1ImageProjectsgetResponseStatusEnum = "canceled"
	V1ImageProjectsgetResponseStatusEnumComplete  V1ImageProjectsgetResponseStatusEnum = "complete"
	V1ImageProjectsgetResponseStatusEnumDraft     V1ImageProjectsgetResponseStatusEnum = "draft"
	V1ImageProjectsgetResponseStatusEnumError     V1ImageProjectsgetResponseStatusEnum = "error"
	V1ImageProjectsgetResponseStatusEnumQueued    V1ImageProjectsgetResponseStatusEnum = "queued"
	V1ImageProjectsgetResponseStatusEnumRendering V1ImageProjectsgetResponseStatusEnum = "rendering"
)
