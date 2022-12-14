package grid

import "fmt"

type Grid struct {
	grid [][]uint8
}

type cell struct {
	x, y int
}

func InitializeGrid(cols, rows int) Grid {
	grid := make([][]uint8, rows)
	for i := range grid {
		grid[i] = make([]uint8, cols)
	}

	return Grid{
		grid: grid,
	}
}

func InitializeSeedGrid(cols, rows int, seedCells []cell) Grid {
	grid := InitializeGrid(cols, rows)
	for _, cell := range seedCells {
		grid.grid[cell.y][cell.x] = 1
	}
	return grid
}

func (grid Grid) Print() {
	for _, i := range grid.grid {
		fmt.Println(i)
	}
}

func (grid Grid) Tick(seedGrid Grid) {
	for y, _ := range grid.grid {
		for x, _ := range grid.grid[y] {
			grid.updateCellStatus(x, y, seedGrid)
		}
	}
}

func (grid Grid) updateCellStatus(x, y int, seedGrid Grid) {
	activeNeighbours := 0
	for c := -1; c <= 1; c++ {
		for r := -1; r <= 1; r++ {
			neighbourColumn := c + x
			neighbourRow := r + y
			if neighbourColumn >= 0 && neighbourRow >= 0 && neighbourColumn < len(seedGrid.grid) && neighbourRow < len(seedGrid.grid[0]) {
				if seedGrid.grid[neighbourRow][neighbourColumn] == 1 {
					activeNeighbours += 1
				}
			}
		}
	}
	if seedGrid.grid[y][x] == 1 {
		activeNeighbours -= 1
	}
	if activeNeighbours < 2 || activeNeighbours > 3 {
		grid.grid[y][x] = 0
	}
	if activeNeighbours == 2 || activeNeighbours == 3 {
		grid.grid[y][x] = seedGrid.grid[y][x]
	}
	if activeNeighbours == 3 {
		grid.grid[y][x] = 1
	}
}
