package utils

import (
	"flag"
	"fmt"
)

func ReadFlags() ([3]int, error) {
	var ms = flag.Int("delay-ms", 2500, "Set the animation speed in milliseconds. Default is 2500 milliseconds")
	var isVerbose = flag.Bool("verbose", false, "Display detailed information about the simulation, including grid size, number of ticks, speed, and map name")
	var isFile = flag.Bool("file", false, "Load the initial grid from a specified file")
	var isEdgesPortal = flag.Bool("edges-portal", false, "Enable portal edges where cells that exit the grid appear on the opposite side")
	var isFullScreen = flag.Bool("fullscreen", false, "Adjust the grid to fit the terminal size with empty cells")
	var isFootPrints = flag.Bool("footprints", false, "Add traces of visited cells, displayed as '∘'")
	var isColored = flag.Bool("colored", false, "Add color to live cells and traces if footprints are enabled")
	flag.Parse()
	fmt.Println(*ms)
	fmt.Println(*isEdgesPortal)
	fmt.Println(*isFullScreen)
	fmt.Println(*isColored)
	fmt.Println(*isFile)
	fmt.Println(*isVerbose)
	fmt.Println(*isFootPrints)

	flags := [3]int{-1, -1, -1}
	//args := os.Args[1:]
	//if len(args) == 0 {
	//	return [3]int{-1, -1, -1}, nil
	//} else {
	//	for i, arg := range args {
	//		if arg == "--help" && i == 0 {
	//			fmt.Println("Usage: go run main.go [options]\n\nOptions:\n  --help        : Show the help message and exit\n  --verbose     : Display detailed information about the simulation, including grid size, number of ticks, speed, and map name\n  --delay-ms=X: Set the animation speed in milliseconds. Default is 2500 milliseconds")
	//			return [3]int{-1, -1, 1}, nil
	//		} else if arg == "--verbose" {
	//			flags[0] = 1
	//		} else if len(arg) > 10 && arg[:11] == "--delay-ms=" {
	//			ms, err := strconv.Atoi(arg[11:])
	//			if err == nil {
	//				flags[1] = ms
	//			} else {
	//				return [3]int{-1, -1, -1}, fmt.Errorf("invalid value for --delay-ms flag: %d", ms)
	//			}
	//		} else {
	//			return [3]int{-1, -1, -1}, fmt.Errorf("no correct flags provided. Use --help for more information")
	//		}
	//	}
	//}
	return flags, nil
}

func flagHelp() {

}

func flagVerbose() {

}

func flagDelayMs() {
}
