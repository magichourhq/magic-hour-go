package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// Style of the face swap video.
type V1FaceSwapCreateBodyStyle struct {
	// * `v1` - May preserve skin detail and texture better, but weaker identity preservation.
	// * `v2` - Faster, sharper, better handling of hair and glasses. stronger identity preservation. (Recommended)
	// * `default` - Use the version we recommend, which will change over time. This is recommended unless you need a specific earlier version. This is the default behavior.
	Version nullable.Nullable[V1FaceSwapCreateBodyStyleVersionEnum] `json:"version,omitempty"`
}
