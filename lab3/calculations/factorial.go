package calculations

import (
	"fmt"
)

func Calculate(n int64, log bool) int64 {
	if log {
		fmt.Println("Before calculations:")
		fmt.Println("Start calculations...")
		fmt.Printf("Calculate %d!\n", n)
	}

	result := int64(1)
	for i := int64(1); i <= n; i++ {
		result *= i
	}

	if log {
		fmt.Println("Calculations complete!")
	}
	fmt.Printf("Factorial of %d is %d\n", n, result)

	return result
}
