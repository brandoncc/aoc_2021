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

func GetCommands() []CommandSet {
  _, filename, _, ok := runtime.Caller(1)

  if !ok {
    fmt.Println(errors.New("There was a problem getting the path to the current file"))
    os.Exit(1)
  }

  file_path := path.Join(path.Dir(filename), "..", "inputs.txt")
  file, err := os.Open(file_path)

  defer file.Close()

  commands := []CommandSet{}

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    command, err := parse_command(scanner.Text())

    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    commands = append(commands, command)
  }

  return commands
}
