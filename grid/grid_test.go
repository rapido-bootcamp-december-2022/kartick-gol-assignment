package grid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGrid(t *testing.T) {
	t.Run("should return a grid of 10 x 10", func(t *testing.T) {
		expectedResponse := [][]uint8{
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		}
		assert.Equal(t, expectedResponse, InitializeGrid(3, 3).grid)
	})
}
