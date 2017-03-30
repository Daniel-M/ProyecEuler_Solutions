// [PROBLEM ID2](https://projecteuler.net/problem=2)
//
// Each new term in the Fibonacci sequence is generated by adding the previous
// two terms. By starting with 1 and 2, the first 10 terms will be:
//
// 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
//
// By considering the terms in the Fibonacci sequence whose values do not
// exceed four million, find the sum of the even-valued terms.
package id2

import (
	"fmt"
	"github.com/Daniel-M/ProjectEuler_Solutions/utilities/fibonacciNumbers"
)

//func main(){
func Solution() {
	fmt.Println("Solution to the problem id2")
	var sum_tot int = 0
	var counter int = 0
	var iBuffer int = 0
	const limit = 4000000

	for i := 1; i < limit; i++ {

		iBuffer = fibonacciNumbers.FibonacciN(i)

		if iBuffer%2 == 0 && iBuffer < limit {
			counter++
			sum_tot += iBuffer
		}
		if iBuffer >= limit {
			break
		}
	}

	fmt.Println("The sum of the first ", counter, "Fibonacci numbers is", sum_tot)
}
