package utils

import (
	"crunch03/globals"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ReadFlags() bool {
	var ms = flag.Int("delay-ms", 2500, "Set the animation speed in milliseconds. Default is 2500 milliseconds")
	var isVerbose = flag.Bool("verbose", false, "Display detailed information about the simulation, including grid size, number of ticks, speed, and map name")
	var isFile = flag.Bool("file", false, "Load the initial grid from a specified file")
	var isEdgesPortal = flag.Bool("edges-portal", false, "Enable portal edges where cells that exit the grid appear on the opposite side")
	var isFullScreen = flag.Bool("fullscreen", false, "Adjust the grid to fit the terminal size with empty cells")
	var isFootPrints = flag.Bool("footprints", false, "Add traces of visited cells, displayed as '∘'")
	var isColored = flag.Bool("colored", false, "Add color to live cells and traces if footprints are enabled")
	var randomCord = flag.String("random", "", "Generate a random grid of the specified width (W) and height (H), min size 3x3")

	var isValid = true
	flag.Parse()

	flagDelayMs(*ms, &isValid)
	flagVerbose(*isVerbose)
	flagEdgesPortal(*isEdgesPortal)
	flagRandom(*randomCord, &isValid)
	flagFullScreen(*isFullScreen)
	flagColored(*isColored)

	fmt.Println(*isFullScreen)
	fmt.Println(*isColored)
	fmt.Println(*isFile)
	fmt.Println(*isFootPrints)

	return isValid
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
		fmt.Println("Invalid delay ms, please enter positive integer.\nSee --help for more information.")
		return
	}
	globals.Interval = time.Duration(ms) * time.Millisecond
}

func flagEdgesPortal(isEdgesPortal bool) {
	globals.IsEdgePortal = isEdgesPortal
}

func flagFullScreen(isFullScreen bool) {
	globals.IsFullScreen = isFullScreen
}

func flagRandom(random string, isValid *bool) {
	coordinates := strings.Split(random, "x")
	if random == "" {
		return
	}
	if len(coordinates) != 2 {
		fmt.Println("Invalid format, provide two numbers in x between them.\nSee --help for more information.")
		*isValid = false
		return
	}
	x, err1 := strconv.Atoi(coordinates[0])
	y, er2 := strconv.Atoi(coordinates[1])
	if err1 != nil || er2 != nil || x <= 2 || y <= 2 {
		fmt.Println("Random flag need two numbers and higher than 2.\nSee --help for more information.")
		*isValid = false
		return
	}
	globals.RandomX = x
	globals.RandomY = y
}

func flagColored(isColored bool) {
	globals.IsColored = isColored
}
