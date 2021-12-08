package solution1

import (
	"math"
	"sort"
)

func Result(inputs []int) int {
  sort.Ints(inputs)

  middle := float64(len(inputs)) / 2
  var target int

  if math.Mod(middle, 2) == 0 {
    target = (inputs[int(middle)] + inputs[int(middle) - 1]) / 2
  }

  target = inputs[int(middle)]

  sum := 0

  for _, v := range inputs {
    sum += int(math.Abs(float64(v - target)))
  }

  return sum
}
