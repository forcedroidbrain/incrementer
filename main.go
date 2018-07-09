package main

import (
	"fmt"

	"github.com/forcedroidbrain/incrementer/incrementer"
)

func main() {
	// Init
	i := incrementer.New()
	fmt.Println(i.GetNumber())
	// Increment
	i.IncrementNumber()
	fmt.Println(i.GetNumber())
	i.IncrementNumber()
	fmt.Println(i.GetNumber())
	// SetMax
	i.SetMaxValue(3)
	i.IncrementNumber()
	fmt.Println(i.GetNumber())
	// Result
	i.IncrementNumber()
	fmt.Println(i.GetNumber())
}
