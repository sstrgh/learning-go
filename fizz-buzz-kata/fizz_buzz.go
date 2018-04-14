package fizzbuzz

import (
  "fmt"
  "strconv"
  "strings"
)

// FizzBuzz takes one integer and iterates incrementally towards that number and adds into the string
//
// (1) "Fizz", if the current integer is divisible by 3 or has the number 3
//
// (2) "Buzz", if the current integer is divisble by 5 or has the number 5
//
// (3) "FizzBuzz", if the current integer has either properties of both point 1 and 2
//
// (4) Otherwise it will just add the integer
func FizzBuzz(numbers int) (output string) {
  hasFizzOrBuzz := false

  for i := 1; i <= numbers; i++ {
      hasFizzOrBuzz = false

      if isDivisbleBy(i, 3) || hasNum(i, 3) {
        output += "Fizz"
        hasFizzOrBuzz = true
      }

      if isDivisbleBy(i, 5) || hasNum(i, 5) {
        output += "Buzz"
        hasFizzOrBuzz = true
      }

      if !hasFizzOrBuzz {
        output += strconv.Itoa(i)
      }

      if i != numbers {
        output += ", "
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
