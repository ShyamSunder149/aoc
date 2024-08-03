
package main

import (
  "fmt"
  "os"
  "log"
  "bufio"
  "strings"
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
  mappings := map[string] string {
    "one" : "1",
    "two" : "2",
    "three" : "3",
    "four" : "4",
    "five" : "5",
    "six" : "6",
    "seven" : "7",
    "eight" : "8",
    "nine" : "9",
    "zero" : "0",
  }
  
  for _, value := range lines {
    
    numeric_string := ""

    flag := true
    i := 0
    ts := ""
    for flag {
      fmt.Println(ts)
      if _, err := strconv.Atoi(string(value[i])); err == nil {
        ts = ts[1:]
        if _, ok := mappings[ts];  ok{
          value = strings.Replace(value, ts, mappings[ts], -1)
          flag = false 
        } else { ts = ts[1:] }
      } else {
        ts += string(value[i])
        if _, ok := mappings[ts]; ok {
        value = strings.Replace(value, ts, mappings[ts], -1)
          flag = false
        }
        i += 1
      }
    } 

    for k, v := range mappings {
      value = strings.Replace(value, k, v, -1)
    }
    
    fmt.Println(value)  
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
