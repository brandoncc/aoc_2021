package day6

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type fish int

func parse_fishes(inputs []string) []fish {
  m := regexp.MustCompile("[0-9]")
  var fishes[]fish

  days := m.FindAllString(inputs[0], -1)

  for _, d := range days {
    num, err := strconv.Atoi(d)

    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    fishes = append(fishes, fish(num))
  }

  return fishes
}

func pass_day(fishes []fish) []fish {
  for i, f := range fishes {
    give_birth := f == 0

    if give_birth {
      fishes[i] = 6
      fishes = append(fishes, fish(8))
    } else {
      fishes[i] = f - 1
    }
  }

  return fishes
}

func solution1(fishes []fish) {
  for i := 0; i < 80; i++ {
    fishes = pass_day(fishes)
  }

  fmt.Println(fmt.Sprintf("The first solution is %d", len(fishes)))
}

func solution2(fishes []fish) {
  for i := 0; i < 256; i++ {
    fishes = pass_day(fishes)
  }

  fmt.Println(fmt.Sprintf("The first solution is %d", len(fishes)))
}

func Run(inputs []string) {
  fishes := parse_fishes(inputs)

  solution1(fishes)
  solution2(fishes)
}
