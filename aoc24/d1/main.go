package main 

import (
  "fmt"
  "bufio"
  "log"
  "os"
  "strings"
  "strconv"
  "slices"
  "math"
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
 
  var firstElements []int 
  var secondElements []int 

  for _, element := range lines {
    nos := strings.Fields(element)
    firstElement, err := strconv.Atoi(nos[0])
    if err != nil {
      panic(err)
    } else {
      firstElements = append(firstElements, firstElement)
    }
    secondElement, err := strconv.Atoi(nos[1])
    if err != nil {
      panic(err)
    } else {
      secondElements = append(secondElements, secondElement)
    }
  }
  
  slices.Sort(firstElements)
  slices.Sort(secondElements)

  difference := 0 

  for index, _ := range firstElements {
   difference += int(math.Abs(float64(firstElements[index] - secondElements[index]))) 
  }
  
  fmt.Println(difference)
}
