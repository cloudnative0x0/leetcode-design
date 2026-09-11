package _3242_neighbor_sum_service

type NeighborSum struct {
	grid   [][]int
	coords map[int][2]int
}

func Constructor(grid [][]int) NeighborSum {
	coords := make(map[int][2]int)

	for i := 0; i < len(grid); i++ {
		for v := 0; v < len(grid[i]); v++ {
			coords[grid[i][v]] = [2]int{i, v}
		}
	}

	return NeighborSum{
		grid:   grid,
		coords: coords,
	}
}

func (ns *NeighborSum) AdjacentSum(value int) int {
	pos, exists := ns.coords[value]
	if !exists {
		return 0
	}

	row, column := pos[0], pos[1]
	sum := 0
	n := len(ns.grid)

	fill := [4][2]int{
		{-1, 0}, // up
		{1, 0},  // down
		{0, -1}, // left
		{0, 1},  // right
	}

	for _, v := range fill {
		nextRow, nextColumn := row+v[0], column+v[1]

		if nextRow >= 0 && nextRow < n && nextColumn >= 0 && nextColumn < n {
			sum += ns.grid[nextRow][nextColumn]
		}
	}

	return sum
}

func (ns *NeighborSum) DiagonalSum(value int) int {
	pos, exists := ns.coords[value]
	if !exists {
		return 0
	}

	row, column := pos[0], pos[1]
	sum := 0
	n := len(ns.grid)

	fill := [4][2]int{
		{-1, -1}, // up-left
		{-1, 1},  // up-right
		{1, -1},  // down-left
		{1, 1},   // down-right
	}

	for _, v := range fill {
		nextRow, nextColumn := row+v[0], column+v[1]

		if nextRow >= 0 && nextRow < n && nextColumn >= 0 && nextColumn < n {
			sum += ns.grid[nextRow][nextColumn]
		}
	}

	return sum
}
