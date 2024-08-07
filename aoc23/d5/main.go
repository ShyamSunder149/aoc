package main

import (
  "fmt"
  "os"
  "bufio"
  "log"
  "strings"
)

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

  seeds := strings.Split(lines[0], " ")[1:]
  fmt.Println(seeds) 

  nos := make(map[int]int)
  for i:=1; i<100; i++ { nos[i] = i }
  fmt.Println(nos)
  
  idx := 1 

  spaces := strings.Split(lines[3]," ")
   

  /* for idx < len(lines) {
    if lines[idx] == "" { 
      idx+=2
    } else {
      
    }
  }*/ 
}
