package day4

import (
	"fmt"
	"os"
	"strconv"
)

func board_index(s []playable, b playable) int {
  for i, v := range s {
    if v == b {
      return i
    }
  }

  return -1
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

func array_has(a []int, v int) bool {
  for _, av := range a {
    if (av == v) {
      return true
    }
  }

  return false
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
