package utils

import (
	"crunch03/globals"
	"flag"
	"fmt"
	"time"
)

func ReadFlags() {
	var ms = flag.Int("delay-ms", 2500, "Set the animation speed in milliseconds. Default is 2500 milliseconds")
	var isVerbose = flag.Bool("verbose", false, "Display detailed information about the simulation, including grid size, number of ticks, speed, and map name")
	var isFile = flag.Bool("file", false, "Load the initial grid from a specified file")
	var isEdgesPortal = flag.Bool("edges-portal", false, "Enable portal edges where cells that exit the grid appear on the opposite side")
	var isFullScreen = flag.Bool("fullscreen", false, "Adjust the grid to fit the terminal size with empty cells")
	var isFootPrints = flag.Bool("footprints", false, "Add traces of visited cells, displayed as '∘'")
	var isColored = flag.Bool("colored", false, "Add color to live cells and traces if footprints are enabled")

	var isValid = true
	flag.Parse()

	flagDelayMs(*ms, &isValid)
	flagVerbose(*isVerbose)
	fmt.Println(*ms)
	fmt.Println(*isEdgesPortal)
	fmt.Println(*isFullScreen)
	fmt.Println(*isColored)
	fmt.Println(*isFile)
	fmt.Println(*isVerbose)
	fmt.Println(*isFootPrints)

	//flags := [3]int{-1, -1, -1}

}

func flagHelp() {

}

func flagVerbose(isVerbose bool) {
	globals.IsVerbose = isVerbose
}

func flagDelayMs(ms int, isValid *bool) {

	if ms <= 0 {
		*isValid = false
		return
	}
	globals.Interval = time.Duration(ms) * time.Millisecond
}
