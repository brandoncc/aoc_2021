package day5

type matrix [][]int

func (m matrix) add_line(l line) {
	point := [2]int{l.x1, l.y1}
	x_v := l.x_velocity()
	y_v := l.y_velocity()

	for {
		if (x_v == -1 && point[0] < l.x2) || (x_v == 1 && point[0] > l.x2) || (y_v == -1 && point[1] < l.y2) || (y_v == 1 && point[1] > l.y2) {
			break
		}

    m[point[1]][point[0]]++
    point[0] = point[0] + x_v
    point[1] = point[1] + y_v
	}
}

func (m matrix) add_lines(lines []line, allow_diagonals bool) {
	for _, l := range lines {
		if allow_diagonals || l.x1 == l.x2 || l.y1 == l.y2 {
			m.add_line(l)
		}
	}
}

func (m matrix) overlap_count() int {
  var count int

  for _, r := range m {
    for _, c := range r {
      if c > 1 {
        count++
      }
    }
  }

  return count
}
