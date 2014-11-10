package main

import (
	"fmt"
	"time"
	"math/rand"
	"math"
)
func main() {
	fmt.Println("Welcome")
	fmt.Println("The time is ", time.Now())
	fmt.Println("A random number ",rand.Intn(10))
	fmt.Println("A random number ",rand.Seed)
	fmt.Println("Now you have %g problems", math.Nextafter(2, 3))
}