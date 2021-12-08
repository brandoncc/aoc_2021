package day7

import (
	"aoc/day7/solution1"
	"aoc/day7/solution2"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func parse_input(input string) []int {
  m := regexp.MustCompile("[0-9]+")
  matches := m.FindAllString(input, -1)
  parsed := []int{}

  for _, v := range matches {
    num, err := strconv.Atoi(v)

    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    parsed = append(parsed, num)
  }

  return parsed
}

func sum(positions []int) int {
  var s int

  for _, v := range positions {
    s += v
  }

  return s
}

func Run(inputs []string) {
  parsed1 := parse_input(inputs[0])
  parsed2 := parse_input(inputs[0])
  
  fmt.Printf("The first solution is %d\n", solution1.Result(parsed1))
  fmt.Printf("The second solution is %d\n", solution2.Result(parsed2))
}
