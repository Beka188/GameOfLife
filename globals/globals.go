package globals

import "time"

var (
	Interval     time.Duration = time.Millisecond * 2500
	IsVerbose    bool          = false
	IsEdgePortal bool          = false
	RandomX      int           = -1
	RandomY      int           = -1
)
