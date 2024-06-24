package utils

import (
	"crunch03/globals"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func readFlags() bool {
	isValid := true

	// Использование флагов Имя/базовое значение/описание
	ms := flag.Int("delay-ms", 2500, "Set the animation speed in milliseconds. Default is 2500 milliseconds")
	isVerbose := flag.Bool("verbose", false, "Display detailed information about the simulation, including grid size, number of ticks, speed, and map name")
	isEdgesPortal := flag.Bool("edges-portal", false, "Enable portal edges where cells that exit the grid appear on the opposite side")
	isFullScreen := flag.Bool("fullscreen", false, "Adjust the grid to fit the terminal size with empty cells")
	isFootPrints := flag.Bool("footprints", false, "Add traces of visited cells, displayed as '∘'")
	isColored := flag.Bool("colored", false, "Add color to live cells and traces if footprints are enabled")
	randomCord := flag.String("random", "", "Generate a random grid of the specified width (W) and height (H), min size 3x3")
	file := flag.String("file", "", "Load the initial grid from a specified file")

	// Разбирает все зарегистрированные флаги командной строки
	flag.Parse()

	// Перебор аргументов командной строки
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--random") || strings.HasPrefix(arg, "-random") {
			flagRandom(*randomCord, &isValid)
		} else if strings.HasPrefix(arg, "-file") || strings.HasPrefix(arg, "--file") {
			flagFile(*file)
		} else if strings.HasPrefix(arg, "-verbose") || strings.HasPrefix(arg, "--verbose") {
			flagVerbose(*isVerbose)
		} else if strings.HasPrefix(arg, "-fullscreen") || strings.HasPrefix(arg, "--fullscreen") {
			flagFullScreen(*isFullScreen)
		}
	}

	// Вызов функций для всех флагов
	flagDelayMs(*ms, &isValid)
	flagVerbose(*isVerbose)
	flagRandom(*randomCord, &isValid)
	flagEdgesPortal(*isEdgesPortal)
	flagFullScreen(*isFullScreen)
	flagColored(*isColored)
	flagFootPrints(*isFootPrints)
	flagFile(*file)

	return isValid
}

func flagHelp() {
}

func flagVerbose(isVerbose bool) {
	if globals.IsFullScreen {
		return
	}
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
	if globals.IsVerbose {
		return
	}
	globals.IsFullScreen = isFullScreen
}

func flagRandom(random string, isValid *bool) {
	coordinates := strings.Split(random, "x")
	if random == "" || globals.FileName != "" {
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
	globals.IsRandom = true
	globals.RandomX = x
	globals.RandomY = y
}

func flagColored(isColored bool) {
	globals.IsColored = isColored
}

func flagFootPrints(isFootPrints bool) {
	globals.IsFootPrint = isFootPrints
}

func flagFile(file string) {
	if file == "" || globals.IsRandom {
		return
	}
	globals.FileName = file
}
