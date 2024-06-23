

# Game of Life

Welcome to the ```Game of Life!``` This is a terminal-based implementation of Conway's Game of Life written in Go. The Game of Life is a cellular automaton devised by the British mathematician John Horton Conway in 1970.
- --
## Run the program
#### Clone the repository
```bash 
git clone https://github.com/Beka188/GameOfLife
cd GameOfLife
```
#### Build the project
```bash 
go build -o game_of_life
```

#### Run the executable
```bash 
./game_of_life
```
- --
## Features
### Add following flags to customize your game

- __--verbose:__  
 Display detailed information about the simulation, including grid size, number of ticks, speed, and map name
- __--delay-ms=X:__  
 Set the animation speed in milliseconds. Default is 2500 milliseconds
- __--file=X:__  
Load the initial grid from a specified file
- __--edges-portal:__  
 Enable portal edges where cells that exit the grid appear on the opposite side
- __--random=WxH:__  
Generate a random grid of the specified width (W) and height (H)
- __--fullscreen:__  
Adjust the grid to fit the terminal size with empty cells
- __--footprints:__  
Add traces of visited cells, displayed as '∘'
- __--colored:__ 
- Add color to live cells and traces if footprints are enabled

### Note: 
Flags (`--random` and `--file`), (`--verbose` and `--fullscreen`) are conflicting, which means whichever comes first will work, the other ignored.
