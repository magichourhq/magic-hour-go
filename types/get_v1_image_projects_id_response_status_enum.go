package types

// The status of the image.
type GetV1ImageProjectsIdResponseStatusEnum string

const (
	GetV1ImageProjectsIdResponseStatusEnumCanceled  GetV1ImageProjectsIdResponseStatusEnum = "canceled"
	GetV1ImageProjectsIdResponseStatusEnumComplete  GetV1ImageProjectsIdResponseStatusEnum = "complete"
	GetV1ImageProjectsIdResponseStatusEnumDraft     GetV1ImageProjectsIdResponseStatusEnum = "draft"
	GetV1ImageProjectsIdResponseStatusEnumError     GetV1ImageProjectsIdResponseStatusEnum = "error"
	GetV1ImageProjectsIdResponseStatusEnumQueued    GetV1ImageProjectsIdResponseStatusEnum = "queued"
	GetV1ImageProjectsIdResponseStatusEnumRendering GetV1ImageProjectsIdResponseStatusEnum = "rendering"
)
