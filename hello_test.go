package hello

import (
	"testing"
	"fmt"
)

func TestHello(t *testing.T) {
	expected := "Hello Go!"
	actual := Hello("Go")
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got: '%s'", expected, actual)
	}
}

func TestSum(t *testing.T) {
	expected := 15
	actual := Sum([]int{1,2,3,4,5})
	if actual != expected {
		t.Errorf("Test failed, expected: %v, got: %v", expected, actual)
	}
}

func ExampleSum() {
	numbers := []int{5,5,5}
	fmt.Println(Sum(numbers))
	// Output:
	// 15
}

// To generate documentation:
// Run godoc: godoc -http=:6060
// godoc -html github.com/antonimassomola/go-tdd > docs.html