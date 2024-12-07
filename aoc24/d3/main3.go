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
    fmt.Println(line)
    m := regexp.MustCompile("mul\\([0-9]{1,},[0-9]{1,}\\)")
    exps := m.FindAllStringSubmatch(line, -1)
    for _, elements := range exps {
      element := elements[0]
      i := 4
      numberStr1 := ""
      for string(element[i]) != "," {
        numberStr1 += string(element[i])
        i += 1
      }
      
      i += 1 
      numberStr2 := ""
      for string(element[i]) != ")" {
        numberStr2 += string(element[i])
        i +=1 
      }

      number1, err := strconv.Atoi(numberStr1)
      if err != nil {
        fmt.Printf("Error while parsing number 1")
      }

      number2, err := strconv.Atoi(numberStr2) 
      if err != nil {
        fmt.Printf("Error while parsing number 2")
      }

      ans += number1 * number2
    }
  }

  fmt.Printf("Answer : %d", ans)
  
}
