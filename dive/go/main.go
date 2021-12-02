package main

import (
	"dive/aoc"
	"dive/submarine1"
	"dive/submarine2"
	"fmt"
)

func solution1(commands []aoc.CommandSet) {
  sub := submarine1.Submarine{}

  for _, v := range commands {
    sub.ProcessCommand(v)
  }

  fmt.Println(fmt.Sprintf("The first answer is %d", sub.Answer()))
}

func solution2(commands []aoc.CommandSet) {
  sub := submarine2.Submarine{}

  for _, v := range commands {
    sub.ProcessCommand(v)
  }

  fmt.Println(fmt.Sprintf("The first answer is %d", sub.Answer()))
}

func main() {
  commands := aoc.GetCommands()
  solution1(commands)
  solution2(commands)
}
