package main

import "testing"

func TestFizzBuzz(t *testing.T) {
  got := FizzBuzz(15)
  wanted := "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, Fizz, 14, FizzBuzz"

  if got != wanted {
    t.Errorf("\ng '%s', \nw '%s'", got, wanted)
  }

}
