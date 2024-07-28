package main

import (
  "fmt"
  "os"
  "log"
  "bufio"
  "strconv"
)

func main ()  {
  f, err := os.Open("input.txt")
  if(err != nil) {
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
    numeric_string := ""
    for _, character := range value {
      if _, err := strconv.Atoi(string(character)); err == nil  {
        numeric_string = numeric_string + string(character)
      }
    }
    
    n := len(numeric_string)
    numeric_string = string(numeric_string[0]) + string(numeric_string[n-1])

    number, err := strconv.Atoi(numeric_string)
    if err != nil {
      log.Fatal(err)
    }

    res = res + number
  }

  fmt.Printf("%d\n", res)
}
