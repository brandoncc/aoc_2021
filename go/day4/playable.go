package day4

type playable interface {
  call_number(number int)
  has_won() bool
  won_with_points() int
  win_solution() int
}

