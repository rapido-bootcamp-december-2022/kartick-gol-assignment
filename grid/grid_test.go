package grid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGrid(t *testing.T) {
	t.Run("should return a grid of 3 x 3", func(t *testing.T) {
		expectedResponse := [][]uint8{
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		}
		assert.Equal(t, expectedResponse, InitializeGrid(3, 3).grid)
	})

	t.Run("should return grid of 4 x 4 with 2 alive cells", func(t *testing.T) {
		expectedResponse := [][]uint8{
			{0, 0, 0, 0},
			{1, 1, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		}
		seedCells := []cell{
			cell{
				x: 1,
				y: 1,
			},
			cell{
				x: 0,
				y: 1,
			},
		}
		assert.Equal(t, expectedResponse, InitializeSeedGrid(4, 4, seedCells).grid)
	})

	t.Run("should return grid of 5 x 5 with toad pattern", func(t *testing.T) {
		expectedResponse := [][]uint8{
			{0, 0, 1, 0, 0},
			{0, 1, 0, 0, 1},
			{0, 1, 0, 0, 1},
			{0, 0, 0, 1, 0},
			{0, 0, 0, 0, 0},
		}
		seedCells := []cell{
			cell{
				x: 1,
				y: 1,
			},
			cell{
				x: 2,
				y: 1,
			},
			cell{
				x: 3,
				y: 1,
			},
			cell{
				x: 2,
				y: 2,
			},
			cell{
				x: 3,
				y: 2,
			},
			cell{
				x: 4,
				y: 2,
			},
		}
		grid := InitializeGrid(5, 5)
		seedGrid := InitializeSeedGrid(5, 5, seedCells)
		grid.Tick(seedGrid)
		assert.Equal(t, expectedResponse, grid.grid)
	})

	t.Run("should return grid of 4 x 4 with block pattern", func(t *testing.T) {
		expectedResponse := [][]uint8{
			{0, 0, 0, 0},
			{0, 1, 1, 0},
			{0, 1, 1, 0},
			{0, 0, 0, 0},
		}
		seedCells := []cell{
			cell{
				x: 1,
				y: 1,
			},
			cell{
				x: 2,
				y: 1,
			},
			cell{
				x: 1,
				y: 2,
			},
			cell{
				x: 2,
				y: 2,
			},
		}
		grid := InitializeGrid(4, 4)
		seedGrid := InitializeSeedGrid(4, 4, seedCells)
		grid.Tick(seedGrid)
		assert.Equal(t, expectedResponse, grid.grid)
	})

	t.Run("should return grid of 4 x 4 with blinker pattern", func(t *testing.T) {
		expectedResponse := [][]uint8{
			{0, 1, 0, 0},
			{0, 1, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 0, 0},
		}
		seedCells := []cell{
			cell{
				x: 1,
				y: 1,
			},
			cell{
				x: 0,
				y: 1,
			},
			cell{
				x: 2,
				y: 1,
			},
		}
		grid := InitializeGrid(4, 4)
		seedGrid := InitializeSeedGrid(4, 4, seedCells)
		grid.Tick(seedGrid)
		assert.Equal(t, expectedResponse, grid.grid)
	})
}
