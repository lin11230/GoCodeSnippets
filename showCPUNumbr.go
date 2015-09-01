package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Number of CPU is ", runtime.NumCPU())
}
