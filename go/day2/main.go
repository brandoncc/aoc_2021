package day2

import (
	"aoc/day2/commands"
	"aoc/day2/submarine1"
	"aoc/day2/submarine2"
	"fmt"
)

func solution1(commands []commands.CommandSet) {
  sub := submarine1.Submarine{}

  for _, v := range commands {
    sub.ProcessCommand(v)
  }

  fmt.Println(fmt.Sprintf("The first answer is %d", sub.Answer()))
}

func solution2(commands []commands.CommandSet) {
  sub := submarine2.Submarine{}

  for _, v := range commands {
    sub.ProcessCommand(v)
  }

  fmt.Println(fmt.Sprintf("The first answer is %d", sub.Answer()))
}

func Run(inputs []string) {
  commands := commands.ParseCommands(inputs)
  solution1(commands)
  solution2(commands)
}
