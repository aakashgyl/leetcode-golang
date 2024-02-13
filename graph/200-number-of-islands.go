package main

import "strconv"

func numIslands(grid [][]byte) int {
	count := 0

	visited := make(map[string]bool)

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if exploreIsland(grid, r, c, visited) == true {
				count++
			}
		}
	}
	return count
}

func exploreIsland(grid [][]byte, r, c int, visited map[string]bool) bool {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] == '0' {
		return false
	}

	key := strconv.Itoa(r) + "," + strconv.Itoa(c)
	if _, ok := visited[key]; ok {
		return false
	}

	visited[key] = true
	exploreIsland(grid, r+1, c, visited)
	exploreIsland(grid, r-1, c, visited)
	exploreIsland(grid, r, c+1, visited)
	exploreIsland(grid, r, c-1, visited)

	return true
}
