package aoc

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
)

type CommandSet struct {
  Name string
  Amount int
}

func parse_command(s string) (CommandSet, error) {
  parts := strings.Split(s, " ")

  command := parts[0]
  amount, err := strconv.Atoi(parts[1])

  if err != nil {
    return CommandSet{}, err
  }

  return CommandSet{command, amount}, nil
}
