package main

import "fmt"

type Position struct {
	X int
	Y int
}

var grid [][]int
var checkeds = map[Position]bool{}
var neighbors = [][2]int{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

func findAreaVirus(i, j int) []Position {

	queue := []Position{}
	selecteds := map[Position]bool{
		{i, j}: true,
	}
	queue = append(queue, Position{i, j})
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for _, neighbor := range neighbors {
			ni := node.X + neighbor[0]
			nj := node.Y + neighbor[1]
			if ni >= 0 && ni < len(grid) && nj >= 0 && nj < len(grid[0]) && grid[ni][nj] == 1 && !selecteds[Position{ni, nj}] {
				queue = append(queue, Position{ni, nj})
				selecteds[Position{ni, nj}] = true
				checkeds[Position{ni, nj}] = true
			}
		}
	}
	result := []Position{}
	for key := range selecteds {
		result = append(result, key)
	}
	return result
}
func getIncreaseNextDay(area []Position) []Position {
	selecteds := map[Position]bool{}
	for _, pos := range area {
		for _, neighbor := range neighbors {
			ni := pos.X + neighbor[0]
			nj := pos.Y + neighbor[1]
			if ni >= 0 && ni < len(grid) && nj >= 0 && nj < len(grid[0]) && grid[ni][nj] == 0 && !selecteds[Position{ni, nj}] {
				selecteds[Position{ni, nj}] = true
			}
		}
	}
	result := []Position{}
	for key := range selecteds {
		result = append(result, key)
	}
	return result
}
func solution() int {
	count := 0
	for {
		areaVirus := [][]Position{}
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[0]); j++ {
				if grid[i][j] == 1 && !checkeds[Position{i, j}] {
					areaVirus = append(areaVirus, findAreaVirus(i, j))
				}
			}
		}
		if len(areaVirus) == 0 {
			break
		}
		virusNextDays := [][]Position{}
		increaseMax := 0
		dayIncreaseMax := 0
		for index, area := range areaVirus {
			virusNextDay := getIncreaseNextDay(area)
			if len(virusNextDay) > increaseMax {
				increaseMax = len(virusNextDay)
				dayIncreaseMax = index
			}
			virusNextDays = append(virusNextDays, virusNextDay)
		}
		count += countWall(areaVirus[dayIncreaseMax])
		for _, pos := range areaVirus[dayIncreaseMax] {
			grid[pos.X][pos.Y] = -1
		}
		areaVirus = append(areaVirus[:dayIncreaseMax], areaVirus[dayIncreaseMax+1:]...)
		virusNextDays = append(virusNextDays[:dayIncreaseMax], virusNextDays[dayIncreaseMax+1:]...)
		nextDay(virusNextDays)
		if len(areaVirus) == 0 {
			break
		}
	}
	return count
}

func countWall(position []Position) (count int) {
	count = 0
	for _, pos := range position {
		for _, neighbor := range neighbors {
			ni := pos.X + neighbor[0]
			nj := pos.Y + neighbor[1]
			if ni >= 0 && ni < len(grid) && nj >= 0 && nj < len(grid[0]) && grid[ni][nj] == 0 {
				count++
			}
		}
	}
	return count
}

func nextDay(virusNextDays [][]Position) {
	for _, virusNextDay := range virusNextDays {
		for _, pos := range virusNextDay {
			grid[pos.X][pos.Y] = 1
		}
	}
}

func containVirus(isInfected [][]int) int {
	grid = isInfected
	checkeds = map[Position]bool{}
	result := solution()
	return result
}
func main() {
	fmt.Println(containVirus([][]int{
		{0, 1, 0, 0, 0, 0, 0, 1}, {0, 1, 0, 1, 0, 0, 0, 1}, {0, 0, 0, 0, 0, 0, 0, 1},
	}))
}
