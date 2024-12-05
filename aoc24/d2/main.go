package main 

import (
  "fmt"
  "bufio"
  "log"
  "os"
  "strings"
  "strconv"
  "reflect"
  "sort"
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
    sorted_slice := make([]int, len(report))
    reversed_slice := make([]int, len(report))

    copy(sorted_slice, report)
    copy(reversed_slice, report)

    sort.Ints(sorted_slice)
    sort.Sort(sort.Reverse(sort.IntSlice(reversed_slice)))
    
    inc_or_dec := (reflect.DeepEqual(sorted_slice, report) || reflect.DeepEqual(reversed_slice, report))
    ok := true
    for i := 1; i < len(sorted_slice); i++ {
      if sorted_slice[i-1] == sorted_slice[i] || sorted_slice[i] - sorted_slice[i-1] > 3 {   
        ok = false 
      } 
    }
    
    if inc_or_dec && ok {
      count++
    }
  }  

  fmt.Println(count)
}
