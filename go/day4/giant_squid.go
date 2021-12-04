package day4

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type playable interface {
  call_number(number int)
  has_won() bool
  won_with_points() int
  win_solution() int
}

type rows [][]int
type cols [][]int
type ints []int

type board struct {
  won bool
  rows rows
  cols cols
  draws ints
}

func get_cols(rows rows) cols {
  var cols cols

  for i := range rows {
    if len(cols) - 1 < i {
      cols = append(cols, make([]int, len(rows)))
    }

    for j := range rows {
      cols[i][j] = rows[j][i]
    }
  }

  return cols
}

func new_board(rows rows) *board {
  return &board{
    rows: rows,
    cols: get_cols(rows),
  }
}

func (b board) display() {
  for _, r := range b.cols {
    fmt.Println(fmt.Sprintf("%d %d %d %d %d", r[0], r[1], r[2], r[3], r[4]))
  }
  fmt.Println("")
}

func (b *board) call_number(number int) {
  b.draws = append(b.draws, number)
}

func (b board) has_won() bool {
  return b.has_winning_row() || b.has_winning_column()
}

func (b board) has_winning_row() bool {
  for _, r := range b.rows {
    if subset(r, b.draws) {
      return true
    }
  }

  return false
}

func (b board) has_winning_column() bool {
  for _, c := range b.cols {
    if subset(c, b.draws) {
      return true
    }
  }

  return false
}

func array_has(a []int, v int) bool {
  for _, av := range a {
    if (av == v) {
      return true
    }
  }

  return false
}

func (b board) won_with_points() int {
  points := 0

  for _, r := range b.rows {
    for _, v := range r {
      if !array_has(b.draws, v) {
        points = points + v
      }
    }
  }

  return points
}

func (b board) win_solution() int {
  return b.won_with_points() * b.draws[len(b.draws) - 1]
}

func int_array(a []string) ints {
  ints := make([]int, len(a))

  for i, v := range a {
    val, err := strconv.Atoi(v)

    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    ints[i] = val
  }

  return ints
}

func subset(needles ints, haystack ints) bool {
  for _, needle := range needles {
    found := false

    for _, hay := range haystack {
      if hay == needle {
        found = true
      }
    }

    if !found {
      return false
    }
  }

  return true
}

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

func index(s []playable, b playable) int {
  for i, v := range s {
    if v == b {
      return i
    }
  }

  return -1
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
      if index(winners, b) >= 0 {
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
  fmt.Println(fmt.Sprintf("The first solution is %d", winners[len(winners) - 1].win_solution()))
}
