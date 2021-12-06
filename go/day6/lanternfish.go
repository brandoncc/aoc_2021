package day6

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type fish int

const day_count = 9

func parse_fishes(inputs []string) []fish {
  m := regexp.MustCompile("[0-9]")
  fishes := make([]fish, day_count)

  days := m.FindAllString(inputs[0], -1)

  for _, d := range days {
    num, err := strconv.Atoi(d)

    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    fishes[num]++
  }

  return fishes
}

func pass_day(fishes []fish) []fish {
  new_fishes := make([]fish, day_count)

  new_fishes[6] = fishes[0] + fishes[7]
  new_fishes[8] = fishes[0]

  new_fishes[0] = fishes[1]
  new_fishes[1] = fishes[2]
  new_fishes[2] = fishes[3]
  new_fishes[3] = fishes[4]
  new_fishes[4] = fishes[5]
  new_fishes[5] = fishes[6]
  new_fishes[7] = fishes[8]

  return new_fishes
}

func count(fishes []fish) fish {
  total := fish(0)

  for i := 0; i < day_count; i++ {
    total += fishes[i]
  }

  return total
}

func solution1(fishes []fish) {
  for i := 0; i < 80; i++ {
    fishes = pass_day(fishes)
  }

  fmt.Println(fmt.Sprintf("The first solution is %d", count(fishes)))
}

func solution2(fishes []fish) {
  for i := 0; i < 256; i++ {
    fishes = pass_day(fishes)
  }

  fmt.Println(fmt.Sprintf("The first solution is %d", count(fishes)))
}

func Run(inputs []string) {
  fishes := parse_fishes(inputs)

  solution1(fishes)
  solution2(fishes)
}
