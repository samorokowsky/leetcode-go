package easy

import (
	"fmt"
	"strconv"
)

// Problem202 (https://leetcode.com/problems/happy-number/description/)
/*
Write an algorithm to determine if a number n is happy.
A happy number is a number defined by the following process:

Starting with any positive integer, replace the number by the sum of the squares of its digits.
Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
Those numbers for which this process ends in 1 are happy.
Return true if n is a happy number, and false if not.

Example 1:
Input: n = 19
Output: true
Explanation:
1^2 + 9^2 = 82
8^2 + 2^2 = 68
6^2 + 8^2 = 100
1^2 + 0^2 + 0^2 = 1

Example 2:
Input: n = 2
Output: false

Constraints:
1 <= n <= 2^31 - 1
*/
func Problem202() {
	var examples = map[int]bool{
		19: true,
		2:  false,
		7:  true,
	}

	for k, v := range examples {
		result := isHappy(k)
		if result != v {
			var log = fmt.Sprintf("Problem202.isHappy - result from '%d' = %d doesn't match expected %d.",
				k, result, v)
			panic(log)
		}
	}

	fmt.Println(" Problem202 - Success passed")

}

func isHappy(n int) bool {

	if n == 1 {
		return true
	}

	var usedNumbers []int
	usedNumbers = append(usedNumbers, n)

	squareSum := n
	var result bool

	for squareSum != 1 {
		squareSum = getSquareSumOfDigits(squareSum)
		if contains(usedNumbers, squareSum) {
			result = false
			break
		} else {
			usedNumbers = append(usedNumbers, squareSum)
		}
		if squareSum == 1 {
			result = true
		}
	}

	return result
}

func getSquareSumOfDigits(n int) int {
	numericString := strconv.Itoa(n)
	var squareSum int
	for i := 0; i < len(numericString); i++ {
		squareSum += int(numericString[i]-'0') * int(numericString[i]-'0')
	}
	return squareSum
}

func contains(slice []int, element int) bool {
	for _, value := range slice {
		if value == element {
			return true
		}
	}
	return false
}
