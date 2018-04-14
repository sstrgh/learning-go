package fizzbuzz

import (
	"fmt"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	got := FizzBuzz(15)
	wanted := "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, Fizz, 14, FizzBuzz"

	if got != wanted {
		t.Errorf("\ng '%s', \nw '%s'", got, wanted)
	}
}

func ExampleFizzBuzz() {
	response := FizzBuzz(5)
	fmt.Println(response)
	// Output: 1, 2, Fizz, 4, Buzz
}

func BenchmarkFizzBuzz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FizzBuzz(15)
	}
}
