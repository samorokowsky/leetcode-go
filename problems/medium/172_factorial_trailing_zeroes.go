package medium

import "fmt"

var counter int

// Problem172 (https://leetcode.com/problems/factorial-trailing-zeroes/description/)
/*
Given an integer n, return the number of trailing zeroes in n!.
Note that n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1.

Example 1:
Input: n = 3
Output: 0
Explanation: 3! = 6, no trailing zero.

Example 2:
Input: n = 5
Output: 1
Explanation: 5! = 120, one trailing zero.

Example 3:
Input: n = 0
Output: 0

Constraints:
0 <= n <= 104

Follow up: Could you write a solution that works in logarithmic time complexity?
*/
func Problem172() {
	var examples = map[int]int{
		3:  0,
		5:  1,
		0:  0,
		30: 7,
	}

	for k, v := range examples {
		result := trailingZeroes(k)
		if result != v {
			var log = fmt.Sprintf("Problem172.trailingZeroes - result from '%d' = %d doesn't match expected %d.",
				k, result, v)
			panic(log)
		}
	}

	fmt.Println(" Problem172 - Success passed")
}

func trailingZeroes(n int) int {
	counterDivisionsBy5 := 0
	for i := 0; i <= n; i = i + 5 {
		if i == 0 {
			continue
		}
		countDivisionsIfDivisible(i, 5)
		counterDivisionsBy5 += counter
		counter = 0
	}

	return counterDivisionsBy5
}

func countDivisionsIfDivisible(num int, divider int) int {
	if num%divider == 0 {
		counter++
		num = countDivisionsIfDivisible(num/divider, divider)
	}

	return num
}
