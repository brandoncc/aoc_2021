package solution2

import (
	"math"
)

func find_min(inputs []int) int {
	min := math.MaxInt

	for _, v := range inputs {
		if v < min {
			min = v
		}
	}

	return min
}

func find_max(inputs []int) int {
	max := -1

	for _, v := range inputs {
		if v > max {
			max = v
		}
	}

	return max
}

func calculate_movements(inputs []int, min, max int) [][]int {
	movements := make([][]int, len(inputs))

	for i := 0; i < len(inputs); i++ {
		starting_location := inputs[i]

		for j := min; j <= max; j++ {
			distance := int(math.Abs(float64(j - starting_location)))
			movements[i] = append(movements[i], distance * (1 + distance) / 2)
		}
	}

	return movements
}

func find_smallest_movement_sum(movements [][]int) int {
  var current int

	movement_possibilities := len(movements[0])
	best := math.MaxInt

	for i := 0; i < movement_possibilities; i++ {
		current = 0

		for j := 0; j < len(movements); j++ {
			current += movements[j][i]
		}

		if current < best {
			best = current
		}
	}

	return best
}

func Result(inputs []int) int {
	min := find_min(inputs)
	max := find_max(inputs)
	movements := calculate_movements(inputs, min, max)
	return find_smallest_movement_sum(movements)
}
