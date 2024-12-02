package main 

import (
  "fmt"
  "bufio"
  "log"
  "os"
  "strings"
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

  var reports [][]int

  for _, element := range lines {
    numbers := strings.Fields(element)
    var report []int
    for _, value := range numbers {
      if convertedValue, err := strconv.Atoi(value); err == nil {
        report = append(report, convertedValue)
      }
    }
    reports = append(reports, report)
  }
 
  count := 0
  for _, report := range reports {
    
    condition := false
    n := len(report)
    isIncreasing := true 
    if report[0] == report[1] {
      continue 
    }
    if report[0] > report[1] {
      isIncreasing = false 
    }

    for i := 1; i < n-1; i++ {
      if report[i] == report[i + 1] {
        condition = false
        break;
      }
      if (isIncreasing && report[i+1] - report[i] < 4 && report[i+1] - report[i] > 0) || (report[i] - report[i + 1] > 0 && !isIncreasing && report[i] - report[i + 1] < 4) {
        condition = true 
      } else {
        condition = false
        break 
      }
    }
    
    if condition {
      count++;
    }
  }

  fmt.Println(count-1)
}
