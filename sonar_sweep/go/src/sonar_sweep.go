package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
)

func get_inputs() ([]int, error) {
  _, current_filename, _, ok := runtime.Caller(1)

  if !ok {
    return []int{}, errors.New("There was a problem identifying the source file path")
  }

  data_file_path := path.Join(path.Dir(current_filename), "..", "..", "inputs.txt")
  file, err := os.Open(data_file_path)

  defer file.Close()

  if err != nil {
    return []int{}, err
  }

  lines := []int{}
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    i, err := strconv.Atoi(scanner.Text())

    if err == nil {
      lines = append(lines, i)
    }
  }

  return lines, nil
}

func deeper_count(inputs []int) int {
  last := inputs[0]
  count := 0

  for _, v := range inputs {
    if v > last {
      count++
    }

    last = v
  }

  return count
}

func main() {
  inputs, err := get_inputs()

  if err != nil {
    fmt.Println(errors.New(fmt.Sprintf("There was an error: %s", err)))
    os.Exit(1)
  }

  fmt.Println(fmt.Sprintf("The answer is %d", deeper_count(inputs)))
}
