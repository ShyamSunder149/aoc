package main

import (
  "fmt"
  "log"
  "bufio"
  "os" 
  "strings"
  "strconv"
)

func max(x, y int) int {
  if x > y {
    return x 
  }
  return y
}

func isAvailable(number int, pros []int) bool {
  for _, value := range pros {
    if value == number { return true }
  }
  return false
}

func pow(x, y int) (r int) {
  r = 1
  for i:=0;i<y;i=i+1 { r  = r * x }
  return
}

func main() {

  f, err := os.Open("input.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer f.Close()

  var lines []string
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

  total := 0
  for _, value := range lines {
    c := -1
    var win []int
    var pro []int

    card_string := strings.Split(value, ":")
    numbers := strings.Split(card_string[1], "|") 
    wins := strings.Split(numbers[0], " ")
    pros := strings.Split(numbers[1], " ")

    for _, number  := range wins {
      nos, err := strconv.Atoi(number)
      if err != nil {
        continue
      }

      win = append(win, nos)
    }

    for _, number := range pros {
      nos, err := strconv.Atoi(number)
      if err != nil {
        continue
      }

      pro = append(pro, nos)
    }
    
    for _, number := range win {
        res := isAvailable(number, pro) 
        if res { c = c + 1 }
    }
    
    if c != -1  { total = total + pow(2, c) }
  }
  
  fmt.Println(total)
}

