package main

import (
	"fmt"
	"math"
)

//An O(1) space and O(n) time complexity function to calculate the nth Fibonacci number
func fibonacci(n int) int64 {
	if n == 0 || n == 1 {
		return int64(n)
	} else {
		var f0, f1 int64
		f0, f1 = 0, 1
		for k := 1; k < n; k++ {
			t := f1
			f1 += f0
			f0 = t
		}
		return f1
	}
}

/*
An O(1) space and O(1) time complexity function that computes the nth Fibonacci formula using
the closed form solution
*/
func fastFibonacci(n int) int64 {
	var a, b float64
	a, b = (1+math.Sqrt(5))/float64(2), (1-math.Sqrt(5))/float64(2)
	return int64((math.Pow(a, float64(n)) - math.Pow(b, float64(n))) / math.Sqrt(5))
}

//A function that returns the array (sequence of Fibonacci numbers) from 0 to n in O(n) time
func fibonacciSequence(n int) []int64 {
	if n == 0 {
		return make([]int64, 1)
	} else if n == 1 {
		fib := make([]int64, 2)
		fib[1] = 1
		return fib
	} else {
		fib := make([]int64, n+1)
		fib[1] = 1
		for k := 2; k <= n; k++ {
			fib[k] = fib[k-1] + fib[k-2]
		}
		return fib
	}

}

func main() {
	fmt.Print("Enter a non-negative integer: ")
	var num int
	_, _ = fmt.Scan(&num)
	fmt.Println("The slow Fibonacci calculator returns", fibonacci(num))
	fmt.Println("The fast Fibonacci calculator returns", fastFibonacci(num))
	fmt.Println("The Fibonacci sequence up to", num, "is", fibonacciSequence(num))
}
