package submarine2

import (
	"aoc/day2/commands"
)

type Submarine struct {
  horizontal int
  depth int
  aim int
}

func (s *Submarine) ProcessCommand(c commands.CommandSet) {
  switch c.Name {
  case "up": s.up(c.Amount)
  case "down": s.down(c.Amount)
  case "forward": s.forward(c.Amount)
  }
}

func (s *Submarine) up(amount int) {
  s.aim -= amount
}

func (s *Submarine) down(amount int) {
  s.aim += amount
}

func (s *Submarine) forward(amount int) {
  s.horizontal += amount
  s.depth += s.aim * amount
}

func (s *Submarine) Answer() int {
  return s.horizontal * s.depth
}
