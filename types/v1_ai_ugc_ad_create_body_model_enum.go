package types

// Defaults to kling-3.0 on paid plans and ltx-2.5 on free plans.
type V1AiUgcAdCreateBodyModelEnum string

const (
	V1AiUgcAdCreateBodyModelEnumKling30 V1AiUgcAdCreateBodyModelEnum = "kling-3.0"
	V1AiUgcAdCreateBodyModelEnumLtx25   V1AiUgcAdCreateBodyModelEnum = "ltx-2.5"
	V1AiUgcAdCreateBodyModelEnumVeo31   V1AiUgcAdCreateBodyModelEnum = "veo3.1"
)
