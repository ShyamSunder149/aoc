package main 

import (
  "fmt"
  "bufio"
  "log"
  "os"
  "regexp"
  "strconv"
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
  
  ans := 0
  for _, line := range lines {

    n := len(line)
    for i := 0; i < n-8; i++ {
      reg, err := regexp.MatchString("^mul\\([0-9]{1,},[0-9]{1,}\\)", line[i:]);
      if err != nil {
        fmt.Printf("Regex is incorrect")
      }

      if reg {

        fmt.Printf("%s\n", line[i:])
        value1str := string(line[i+4])
        j := i + 5
        for string(line[j]) != "," {
          value1str += string(line[j])
          j += 1 
        }

        value1, err := strconv.Atoi(value1str)
        if err != nil {
          fmt.Printf("Error while converting value 1\n")
        }
      
        j += 1 
        value2str := string(line[j])
        j += 1 
        for string(line[j]) != ")" {
          value2str += string(line[j])
          j += 1 
        }

        value2, err := strconv.Atoi(value2str)
        if err != nil {
          fmt.Printf("Error while converting value 2\n")
        }
          
        fmt.Printf("%d : %d -> values\n", value1, value2)

        ans += value1 * value2 
      }
    }

  }

  fmt.Printf("Answer : %d", ans)
  
}
