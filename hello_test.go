package main_test

import (
	"fmt"
	"testing"
)

//This does not seems to be called by "go test" as doc says... - revisit
//https://golang.org/cmd/go/#hdr-GOPATH_environment_variable
func BenchmarkXXX(b *testing.B) {
	fmt.Println("IN BENCHMARK FUNC")
}

//This is automatically called by "go test"
func TestXXX(t *testing.T) {
	fmt.Println("IN TEST FUNC")
}

//Special test case that checks output matching the last comment... (probably not very useful)
//This is automatically called by "go test"
func Example2Println() {
	//test()
	// Output: test
}

func ExamplePrintln() {
	fmt.Println("The output of\nthis example.")
	// Output: The output of
	// this example.
}
