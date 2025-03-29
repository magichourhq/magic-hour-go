package types

// The status of the image.
type V1ImageProjectsGetResponseStatusEnum string

const (
	V1ImageProjectsGetResponseStatusEnumCanceled  V1ImageProjectsGetResponseStatusEnum = "canceled"
	V1ImageProjectsGetResponseStatusEnumComplete  V1ImageProjectsGetResponseStatusEnum = "complete"
	V1ImageProjectsGetResponseStatusEnumDraft     V1ImageProjectsGetResponseStatusEnum = "draft"
	V1ImageProjectsGetResponseStatusEnumError     V1ImageProjectsGetResponseStatusEnum = "error"
	V1ImageProjectsGetResponseStatusEnumQueued    V1ImageProjectsGetResponseStatusEnum = "queued"
	V1ImageProjectsGetResponseStatusEnumRendering V1ImageProjectsGetResponseStatusEnum = "rendering"
)
