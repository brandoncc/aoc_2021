package day4

import "fmt"

type board struct {
  won bool
  rows rows
  cols cols
  draws ints
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
