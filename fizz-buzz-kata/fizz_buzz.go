package main

import (
  "fmt"
  "strconv"
  "strings"
)

func FizzBuzz(numOfStudents int) (answers string) {
  hasFizzOrBuzz := false

  for i := 1; i <= numOfStudents; i++ {
      hasFizzOrBuzz = false

      if isDivisbleBy(i, 3) || hasNum(i, 3) {
        answers += "Fizz"
        hasFizzOrBuzz = true
      }

      if isDivisbleBy(i, 5) || hasNum(i, 5) {
        answers += "Buzz"
        hasFizzOrBuzz = true
      }

      if !hasFizzOrBuzz {
        answers += strconv.Itoa(i)
      }

      if i != numOfStudents {
        answers += ", "
      }

    }
  return
}

func isDivisbleBy(num int, comparisonNum int) bool {
  return num%comparisonNum == 0
}

func hasNum(num int, comparisonNum int) bool {
  return strings.ContainsAny(strconv.Itoa(num), strconv.Itoa(comparisonNum))
}

func main() {
	fmt.Println(FizzBuzz(15))
}
