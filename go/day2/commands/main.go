package commands

import (
	"fmt"
	"os"
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

func ParseCommands(inputs []string) []CommandSet {
  var parsed []CommandSet

  for _, v := range inputs {
    command, err := parse_command(v)

    if err != nil {
      fmt.Println(fmt.Sprintf("There was a problem parsing input %s", v))
      os.Exit(1)
    }

    parsed = append(parsed, command)
  }

  return parsed
}

