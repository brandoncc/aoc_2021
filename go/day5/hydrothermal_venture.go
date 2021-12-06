package day5

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func parse_line(input string) line {
	numbers := regexp.MustCompile("[0-9]+")
	matches := numbers.FindAllString(input, -1)
	l := line{}

	for i, v := range matches {
		n, err := strconv.Atoi(v)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		switch i {
		case 0:
			l.x1 = n
		case 1:
			l.y1 = n
		case 2:
			l.x2 = n
		case 3:
			l.y2 = n
		}
	}

	return l
}

func parse_lines(inputs []string) []line {
	parsed := []line{}

	for _, v := range inputs {
		parsed = append(parsed, parse_line(v))
	}

	return parsed
}

func find_max(lines []line, axis rune) int {
	var max int

	for _, v := range lines {
		if axis == 'x' {
			if v.x1 > max {
				max = v.x1
			}

			if v.x2 > max {
				max = v.x2
			}
		}

		if axis == 'y' {
			if v.y1 > max {
				max = v.y1
			}

			if v.y2 > max {
				max = v.y2
			}
		}
	}

	return max
}

func build_matrix(lines []line) matrix {
	row_count := find_max(lines, 'x')
	col_count := find_max(lines, 'y')
	m := matrix{}

	for i := 0; i <= row_count; i++ {
		row := make([]int, col_count+1)
		m = append(m, row)
	}

	return m
}

func solution1(lines []line) {
	matrix := build_matrix(lines)
	matrix.add_lines(lines, false)
	fmt.Println(fmt.Sprintf("The first solution is %d", matrix.overlap_count()))
}

func solution2(lines []line) {
	matrix := build_matrix(lines)
	matrix.add_lines(lines, true)
	fmt.Println(fmt.Sprintf("The second solution is %d", matrix.overlap_count()))
}

func Run(inputs []string) {
	lines := parse_lines(inputs)

	solution1(lines)
	solution2(lines)
}
