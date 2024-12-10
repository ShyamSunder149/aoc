package main 

import (
  "fmt"
  "bufio"
  "log"
  "os"
)

func main() {
  
  f, err := os.Open("./input.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer f.Close()

  var lines []string
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

  // convert arr to 2d matrix
  var xmas [][]rune 
  for _, line := range lines {
    var xmas_line []rune 
    for _, element := range line {
        xmas_line = append(xmas_line, element)
    }
    xmas = append(xmas, xmas_line)
  } 

  count := 0  
  xmas_len := len(xmas)
  line_len := len(xmas[0])
  fmt.Printf("Xmas Len : %d & Line Len : %d\n", xmas_len, line_len)

  for i, line := range xmas {
    for j, element := range line {
      if element != 'X' {
        continue
      }

        for _, r := range []int{-1, 0, 1} {
          for _, c := range []int{-1, 0, 1} {
            if !(i+r*3 >= 0 && i+r*3 < len(xmas) && j+c*3 >= 0 && j+c*3 < len(line)) {
              continue
            } 
            
            if xmas[i+r][j+c] == 'M' && xmas[i+r*2][j+c*2] == 'A' && xmas[i+r*3][j+c*3] == 'S' {
              count++
            }
          }
        }
    }
  }
  fmt.Printf("%d\n", count)
}
