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

  mappings := make(map[int]int)
  
  for _, element := range firstElements {
    if _, exists := mappings[element]; exists {
      continue
    }

    count := 0
    for _, secondListElement := range secondElements {
      if secondListElement == element {
        count += 1
      }
    }

    if count > 0 {
      mappings[element] = count
    }
  }
  
  var product int 

  for _, element := range firstElements {
    if value, exists := mappings[element]; exists {
      product += element * value
    }
  }

  fmt.Println(product)
}
