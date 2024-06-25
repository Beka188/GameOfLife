package globals

import "time"

var (
	Interval     time.Duration = time.Millisecond * 2500
	IsVerbose    bool          = false
	IsEdgePortal bool          = false
	IsFullScreen bool          = false
	IsColored    bool          = false
	IsFootPrint  bool          = false
	IsRandom     bool          = false
	FileName                   = ""
	RandomX      int           = 0
	RandomY      int           = 0
)

const (
	ColorEmpty     = "\033[90m" // Grey
	ColorLive      = "\033[93m" // Yellow
	ColorFootprint = "\033[95m" // Magenta
	ResetColor     = "\033[0m"
)
