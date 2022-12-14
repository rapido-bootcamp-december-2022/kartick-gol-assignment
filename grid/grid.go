package grid

type Grid struct {
	grid [][]uint8
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
