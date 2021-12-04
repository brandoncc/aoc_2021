package day4

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type rows [][]int
type cols [][]int
type ints []int

func build_boards(inputs []string) []*board {
  i := 0
  boards := []*board{}
  var int_inputs rows

  for i, v := range inputs {
    items := strings.Fields(v)
     
    if len(int_inputs) - 1 < i {
      int_inputs = append(int_inputs, []int{})
    }

    for _, c := range items {
      conv, err := strconv.Atoi(c)

      if err != nil {
        fmt.Println(err)
        os.Exit(1)
      }

      int_inputs[i] = append(int_inputs[i], conv)
    }
  }

  for {
    if i >= len(int_inputs) {
      break
    }

    boards = append(boards, new_board(int_inputs[i:i+5]))
    i = i + 6
  }

  return boards
}

func Run(inputs []string) {
  draws := int_array(strings.Split(inputs[0], ","))
  boards := build_boards(inputs[2:])
  winners := []playable{}

  for _, d := range draws {
    if len(winners) == len(boards) {
      continue
    }

    for _, b := range boards {
      if board_index(winners, b) >= 0 {
        continue
      }

      b.call_number(d)

      if b.has_won() {
        winners = append(winners, b)
      }
    }
  }

  for _, b := range boards {
    b.display()
  }

  fmt.Println(fmt.Sprintf("There are %d winners", len(winners)))
  fmt.Println(fmt.Sprintf("The first solution is %d", winners[0].win_solution()))
  fmt.Println(fmt.Sprintf("The second solution is %d", winners[len(winners) - 1].win_solution()))
}
