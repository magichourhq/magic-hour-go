package types

// The status of the video.
type GetV1VideoProjectsIdResponseStatusEnum string

const (
	GetV1VideoProjectsIdResponseStatusEnumCanceled  GetV1VideoProjectsIdResponseStatusEnum = "canceled"
	GetV1VideoProjectsIdResponseStatusEnumComplete  GetV1VideoProjectsIdResponseStatusEnum = "complete"
	GetV1VideoProjectsIdResponseStatusEnumDraft     GetV1VideoProjectsIdResponseStatusEnum = "draft"
	GetV1VideoProjectsIdResponseStatusEnumError     GetV1VideoProjectsIdResponseStatusEnum = "error"
	GetV1VideoProjectsIdResponseStatusEnumQueued    GetV1VideoProjectsIdResponseStatusEnum = "queued"
	GetV1VideoProjectsIdResponseStatusEnumRendering GetV1VideoProjectsIdResponseStatusEnum = "rendering"
)
