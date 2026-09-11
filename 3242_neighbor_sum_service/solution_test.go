package _3242_neighbor_sum_service

import "testing"

func TestNeighborSum(t *testing.T) {
	grid := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}
	ns := Constructor(grid)

	tests := []struct {
		name     string
		value    int
		adjacent int
		diagonal int
	}{
		{name: "center", value: 4, adjacent: 16, diagonal: 16},
		{name: "top edge", value: 1, adjacent: 6, diagonal: 8},
		{name: "top left corner", value: 0, adjacent: 4, diagonal: 4},
		{name: "bottom right corner", value: 8, adjacent: 12, diagonal: 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ns.AdjacentSum(tt.value); got != tt.adjacent {
				t.Errorf(
					"AdjacentSum(%d) = %d; want %d",
					tt.value,
					got,
					tt.adjacent,
				)
			}

			if got := ns.DiagonalSum(tt.value); got != tt.diagonal {
				t.Errorf(
					"DiagonalSum(%d) = %d; want %d",
					tt.value,
					got,
					tt.diagonal,
				)
			}
		})
	}
}

func TestSingleCell(t *testing.T) {
	ns := Constructor([][]int{{7}})

	if got := ns.AdjacentSum(7); got != 0 {
		t.Errorf("AdjacentSum(7) = %d; want 0", got)
	}

	if got := ns.DiagonalSum(7); got != 0 {
		t.Errorf("DiagonalSum(7) = %d; want 0", got)
	}
}

func TestUnknownValue(t *testing.T) {
	ns := Constructor([][]int{
		{1, 2},
		{3, 4},
	})

	if got := ns.AdjacentSum(100); got != 0 {
		t.Errorf("AdjacentSum(100) = %d; want 0", got)
	}

	if got := ns.DiagonalSum(100); got != 0 {
		t.Errorf("DiagonalSum(100) = %d; want 0", got)
	}
}

func TestAllCellsAgainstDirectCalculation(t *testing.T) {
	grid := [][]int{
		{15, 3, 11, 8},
		{9, 14, 0, 6},
		{4, 1, 13, 10},
		{12, 7, 5, 2},
	}
	ns := Constructor(grid)

	adjacentDirections := [][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	diagonalDirections := [][2]int{
		{-1, -1},
		{-1, 1},
		{1, -1},
		{1, 1},
	}

	for row := range grid {
		for column, value := range grid[row] {
			wantAdjacent := directSum(
				grid,
				row,
				column,
				adjacentDirections,
			)

			wantDiagonal := directSum(
				grid,
				row,
				column,
				diagonalDirections,
			)

			if got := ns.AdjacentSum(value); got != wantAdjacent {
				t.Errorf(
					"AdjacentSum(%d) = %d; want %d",
					value,
					got,
					wantAdjacent,
				)
			}

			if got := ns.DiagonalSum(value); got != wantDiagonal {
				t.Errorf(
					"DiagonalSum(%d) = %d; want %d",
					value,
					got,
					wantDiagonal,
				)
			}
		}
	}
}

func directSum(
	grid [][]int,
	row int,
	column int,
	directions [][2]int,
) int {
	sum := 0

	for _, direction := range directions {
		nextRow := row + direction[0]
		nextColumn := column + direction[1]

		if nextRow >= 0 &&
			nextRow < len(grid) &&
			nextColumn >= 0 &&
			nextColumn < len(grid) {
			sum += grid[nextRow][nextColumn]
		}
	}

	return sum
}
