package globals

import "time"

var (
	Interval     time.Duration = time.Millisecond * 2500
	IsVerbose    bool          = false
	IsEdgePortal bool          = false
	IsFullScreen bool          = false
	IsColored    bool          = false
	RandomX      int           = 0
	RandomY      int           = 0
)
