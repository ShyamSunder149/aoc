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
  
  res := 0

  for _, value := range lines {
    game_id_string := strings.Split(value, ":")
    game_id := strings.Split(game_id_string[0], " ")

    game_strings := strings.Split(game_id_string[1], ";")

    r := 0
    g := 0
    b := 0

    for _, game_string := range game_strings {
      color_strings := strings.Split(game_string, ",")

      for _, color_string := range color_strings {
        count_string := strings.Split(color_string, " ")

        count, err := strconv.Atoi(count_string[1])
        if err != nil {
          log.Fatal(err)
        }

        switch count_string[2] {
          case "red" : r = max(r, count) 
          case "blue" :  b = max(b, count)
          case "green" : g = max(g, count) 
        }
      }
    }

    if r <= 12 && g <= 13 && b <= 14 { 
      game_id_int, err := strconv.Atoi(string(game_id[1]))
      if(err != nil) {
        log.Fatal(err)
      }

      res = res + game_id_int 
    }
  }

  fmt.Println(res)
}
