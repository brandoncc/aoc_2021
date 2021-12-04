package main

import (
	"bufio"
	"errors"
	"path"
	"runtime"
	"aoc/day1"
	"aoc/day2"
	"aoc/day3"

	// "aoc/day4"
	// "aoc/day5"
	// "aoc/day6"
	// "aoc/day7"
	// "aoc/day8"
	// "aoc/day9"
	// "aoc/day10"
	// "aoc/day11"
	// "aoc/day12"
	// "aoc/day13"
	// "aoc/day14"
	// "aoc/day15"
	// "aoc/day16"
	// "aoc/day17"
	// "aoc/day18"
	// "aoc/day19"
	// "aoc/day20"
	// "aoc/day21"
	// "aoc/day22"
	// "aoc/day23"
	// "aoc/day24"
	// "aoc/day25"
	"fmt"
	"os"
	"strconv"
)

func get_commands(day int) []string {
  _, filename, _, ok := runtime.Caller(1)

  if !ok {
    fmt.Println(errors.New("There was a problem getting the path to the current file"))
    os.Exit(1)
  }

  file_path := path.Join(path.Dir(filename), "..", "inputs", fmt.Sprintf("day%d.txt", day))
  file, err := os.Open(file_path)

  defer file.Close()

  commands := []string{}

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    commands = append(commands, scanner.Text())
  }

  return commands
}

func run(day int) {
  switch day {
  case 1: day1.Run(get_commands(day))
  case 2: day2.Run(get_commands(day))
  case 3: day3.Run(get_commands(day))
  // case 4: day4.Run(get_commands(day))
  // case 5: day5.Run(get_commands(day))
  // case 6: day6.Run(get_commands(day))
  // case 7: day7.Run(get_commands(day))
  // case 8: day8.Run(get_commands(day))
  // case 9: day9.Run(get_commands(day))
  // case 10: day10.Run(get_commands(day))
  // case 11: day11.Run(get_commands(day))
  // case 12: day12.Run(get_commands(day))
  // case 13: day13.Run(get_commands(day))
  // case 14: day14.Run(get_commands(day))
  // case 15: day15.Run(get_commands(day))
  // case 16: day16.Run(get_commands(day))
  // case 17: day17.Run(get_commands(day))
  // case 18: day18.Run(get_commands(day))
  // case 19: day19.Run(get_commands(day))
  // case 20: day20.Run(get_commands(day))
  // case 21: day21.Run(get_commands(day))
  // case 22: day22.Run(get_commands(day))
  // case 23: day23.Run(get_commands(day))
  // case 24: day24.Run(get_commands(day))
  // case 25: day25.Run(get_commands(day))
  }
}

func main() {
  if len(os.Args) != 2 {
    fmt.Println("You must provide the day number to run")
    os.Exit(1)
  }

  day, err := strconv.Atoi(os.Args[1])

  if err != nil {
    fmt.Println("There was an error parsing the day number that you provided")
    os.Exit(1)
  }

  if day < 1 || day > 25 {
    fmt.Println("Advent of Code is only 25 days long, please enter a number in the range 1-25")
    os.Exit(1)
  }

  run(day)
}
