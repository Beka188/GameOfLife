package utils

import "os"
import "fmt"

func ReadFlags() []string {
	return os.Args[1:]

}

func flagHelp() {

}

func flagVerbose() {

}

func flagDelayMs() {
	fmt.Println("Usage: go run main.go [options]\n\nOptions:\n  --help        : Show the help message and exit\n  --verbose     : Display detailed information about the simulation, including grid size, number of ticks, speed, and map name\n  --delay-ms=X: Set the animation speed in milliseconds. Default is 2500 milliseconds")
}
