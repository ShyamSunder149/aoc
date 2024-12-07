package main 

import (
  "fmt"
  "bufio"
  "log"
  "os"
  "regexp"
  "strconv"
)

func check_mul(start, end int, line string) int {
      fmt.Printf("Checking in the range of %d to %d\n", start, end)
      ans := 0

      for i:= start; i < end - 8; i++ {
      reg, err := regexp.MatchString("^mul\\([0-9]{1,},[0-9]{1,}\\)", line[i:]);
      if err != nil {
        fmt.Printf("Regex is incorrect")
      }

      if reg {

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
          
        ans += value1 * value2 
      }
    }
    return ans
}

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
    
  var do_array []int 
  var dont_array []int 
  n := len(line)

    for i := 0; i < n-4; i++ {
      regmatch, err := regexp.MatchString("^do\\(\\)", line[i:])
      if err != nil {
        fmt.Printf("Error while checking do array")
      }
      
      if regmatch {
      do_array = append(do_array, i)
      continue
      }

      if(i < n-7) {
        regmatch2, err := regexp.MatchString("^don't\\(\\)", line[i:])
        if err != nil {
          fmt.Printf("Error while checking the dont array")
        }
        
        if regmatch2 {
        dont_array = append(dont_array, i)
      }
    }
    }

    fmt.Println(do_array)
    fmt.Println(dont_array)

    ans = check_mul(0, dont_array[0], line)
    fmt.Printf("Answer for 1st string : %d\n", ans)
    do_array_index := 0
    dont_array_index := 1

    for do_array[do_array_index] < dont_array[0] {
      do_array_index += 1 
    } 

    do_array_len := len(do_array)
    dont_array_len := len(dont_array)
    for do_array_index < do_array_len && dont_array_index < dont_array_len {
        fmt.Printf("For loop check : %d - %d\n", do_array[do_array_index], dont_array[dont_array_index])
        for dont_array_index < dont_array_len-1 && dont_array[dont_array_index] < do_array[do_array_index] {
          dont_array_index += 1 
          continue
        }

        fmt.Printf("Range : %d - %d\n", do_array[do_array_index], dont_array[dont_array_index])
        if dont_array_index > dont_array_len {
          break
        }

        ans += check_mul(do_array[do_array_index], dont_array[dont_array_index], line)
        do_array_index += 1 
        dont_array_index += 1
    }

    if do_array[do_array_len - 1] > dont_array[dont_array_len - 1] && do_array_index < do_array_len - 1 {
      fmt.Printf("Final leg : %d - %d\n", do_array[do_array_index], n)
      ans += check_mul(do_array[do_array_len - 1], n, line)
    }
  }
  fmt.Printf("Answer : %d", ans)
  
}
