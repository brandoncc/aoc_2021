package day3

import (
	"fmt"
	"os"
	"strconv"
)

func most_common_bit(inputs []string, position int, default_value rune) rune {
  ones := 0

  for _, v := range inputs {
    if v[position] == '1' {
      ones++
    }
  }

  if float64(ones) == float64(len(inputs)) / 2 {
    return default_value
  }

  if ones > len(inputs) / 2 {
    return '1'
  }

  return '0'
}

func binary_to_int(b string) int64 {
  i, err := strconv.ParseInt(b, 2, 64)

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  return i
}

func solution1(inputs []string) int64 {
  gamma := ""
  epsilon := ""

  for i := 0; i < len(inputs[0]); i++ {
    if most_common_bit(inputs, i, '1') == '1' {
      gamma = gamma + "1"
      epsilon = epsilon + "0"

      continue
    }

    gamma = gamma + "0"
    epsilon = epsilon + "1"
  }

  return binary_to_int(gamma) * binary_to_int(epsilon)
}

func oxygen(inputs []string, comparison_bit int) string {
  if len(inputs) == 1 {
    return inputs[0]
  }

  if comparison_bit > len(inputs[0]) {
    return ""
  }

  var filtered []string
  common_bit := most_common_bit(inputs, comparison_bit, '1')

  for _, v := range inputs {
    if v[comparison_bit] == byte(common_bit) {
      filtered = append(filtered, v)
    }
  }

  return oxygen(filtered, comparison_bit + 1)
}

func c02(inputs []string, comparison_bit int) string {
  if len(inputs) == 1 {
    return inputs[0]
  }

  if comparison_bit > len(inputs[0]) {
    return ""
  }

  var filtered []string
  common_bit := most_common_bit(inputs, comparison_bit, '1')

  for _, v := range inputs {
    if v[comparison_bit] != byte(common_bit) {
      filtered = append(filtered, v)
    }
  }

  return c02(filtered, comparison_bit + 1)
}

func solution2(inputs []string) int64 {
  o := oxygen(inputs, 0)
  c := c02(inputs, 0)

  return binary_to_int(o) * binary_to_int(c)
}

func Run(inputs []string) {
  fmt.Println(fmt.Sprintf("The first solution is %d", solution1(inputs)))
  fmt.Println(fmt.Sprintf("The second solution is %d", solution2(inputs)))
}
