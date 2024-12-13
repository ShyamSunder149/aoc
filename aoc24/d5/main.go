package main 

import (
  "fmt"
  "bufio"
  "log"
  "os"
  "strings"
  "strconv"
  "slices"
)

func convert(s string) int {
  v, err := strconv.Atoi(s)
  if err != nil {
    fmt.Println("Some Error while converting")
  }

  return v 
} 

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

// input parsing 

  space_idx := 0
  for idx, element := range lines {
    if element == "" {
      space_idx = idx 
    }
  } 

  m := make(map[int][]int) 
  n := len(lines)
  value := 0

  for i:= 0;i<space_idx;i++ {
    line := strings.Split(lines[i], "|")
    x := convert(line[0])
    y := convert(line[1])
    
    m[y] = append(m[y], x)
  }

  for i:=space_idx+1;i<n;i++ {
    vs := strings.Split(lines[i], ",")
    var vsi []int 
    for _, e := range vs {
      vsi = append(vsi, convert(e))
    }

    ok := true
    for a, x := range vsi {
      for b, y := range vsi {
        if a < b && slices.Contains(m[x], y) {
          ok = false
        }
      }
    }

    if ok {
      value += vsi[len(vsi)/2]
    }
  }

  fmt.Println(value)
}
