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
      if element != 'A' {
        continue
      }
      
      if !(i-1 >= 0 && i+1 < xmas_len && j-1 >=0 && j+1 < line_len) {
        continue 
      }

      if xmas[i-1][j-1] == 'M' && xmas[i+1][j+1] == 'S' {
        if (xmas[i+1][j-1] == 'M' && xmas[i-1][j+1] == 'S') || (xmas[i+1][j-1] == 'S' && xmas[i-1][j+1] == 'M') {
          count++
        }
      }
      
      if xmas[i-1][j-1] == 'S' && xmas[i+1][j+1] == 'M' {
        if (xmas[i+1][j-1] == 'M' && xmas[i-1][j+1] == 'S') || (xmas[i+1][j-1] == 'S' && xmas[i-1][j+1] == 'M') {
          count++
        }
      }

    }
  }
  fmt.Printf("%d\n", count)
}
