package day5

type line struct {
  x1, x2, y1, y2 int
}

func (l line) x_velocity() int {
  if l.x1 > l.x2 {
    return -1
  }

  if l.x1 < l.x2 {
    return 1
  }

  return 0
}

func (l line) y_velocity() int {
  if l.y1 > l.y2 {
    return -1
  }

  if l.y1 < l.y2 {
    return 1
  }

  return 0
}
