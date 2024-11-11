package medium

import "fmt"

// Problem264 (https://leetcode.com/problems/ugly-number-ii/description/)
/*
An ugly number is a positive integer whose prime factors are limited to 2, 3, and 5.
Given an integer n, return the nth ugly number.

Example 1:
Input: n = 10
Output: 12
Explanation: [1, 2, 3, 4, 5, 6, 8, 9, 10, 12] is the sequence of the first 10 ugly numbers.

Example 2:
Input: n = 1
Output: 1
Explanation: 1 has no prime factors, therefore all of its prime factors are limited to 2, 3, and 5.

Constraints:
1 <= n <= 1690
*/
func Problem264() {
	var examples = map[int]int{
		10:  12,
		1:   1,
		100: 1536,
		200: 16200,
	}

	for k, v := range examples {
		result := nthUglyNumber(k)
		if result != v {
			var log = fmt.Sprintf("Problem264.nthUglyNumber - result from '%d' = %d doesn't match expected %d.",
				k, result, v)
			panic(log)
		}
	}

	fmt.Println(" Problem264 - Success passed")

}

func nthUglyNumber(n int) int {
	var uglyArray []int
	var uglyHashArray []int
	uglyArray = append(uglyArray, 1)

	for len(uglyArray) < n {
		current := uglyArray[len(uglyArray)-1]

		x := current * 2
		y := current * 3
		z := current * 5

		if len(uglyHashArray) != 0 {
			minHashElement := findMin(uglyHashArray)
			if minHashElement < x && minHashElement < y && minHashElement < z {
				uglyArray = append(uglyArray, minHashElement)
				uglyHashArray = removeByValue(uglyHashArray, minHashElement)
				if !contains(uglyHashArray, x) {
					uglyHashArray = append(uglyHashArray, x)
				}
				if !contains(uglyHashArray, y) {
					uglyHashArray = append(uglyHashArray, y)
				}
				if !contains(uglyHashArray, z) {
					uglyHashArray = append(uglyHashArray, z)
				}
			} else if x < minHashElement && x < y && x < z {
				uglyArray = append(uglyArray, x)
				addToSet(uglyHashArray, y)
				addToSet(uglyHashArray, z)
			} else if y < minHashElement && y < x && y < z {
				uglyArray = append(uglyArray, y)
				addToSet(uglyHashArray, x)
				addToSet(uglyHashArray, z)
			} else {
				uglyArray = append(uglyArray, z)
				addToSet(uglyHashArray, x)
				addToSet(uglyHashArray, y)
			}
		} else {
			if x < y && x < z {
				uglyArray = append(uglyArray, x)
				uglyHashArray = append(uglyHashArray, y)
				uglyHashArray = append(uglyHashArray, z)
			} else if y < x && y < z {
				uglyArray = append(uglyArray, y)
				uglyHashArray = append(uglyHashArray, x)
				uglyHashArray = append(uglyHashArray, z)
			} else {
				uglyArray = append(uglyArray, z)
				uglyHashArray = append(uglyHashArray, x)
				uglyHashArray = append(uglyHashArray, y)
			}
		}
	}

	return uglyArray[n-1]
}

func addToSet(slice []int, element int) {
	if !contains(slice, element) {
		slice = append(slice, element)
	}
}

func findMin(slice []int) int {
	if len(slice) == 0 {
		panic("Slice is empty")
	}
	minimum := slice[0]

	for _, value := range slice {
		if value < minimum {
			minimum = value
		}
	}

	return minimum
}

func removeByValue(slice []int, value int) []int {
	for i, v := range slice {
		if v == value {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func contains(slice []int, element int) bool {
	for _, value := range slice {
		if value == element {
			return true
		}
	}
	return false
}

func nthUglyNumberSlowVersion(n int) int {
	integer := 1
	var uglyNumbers []int

	for len(uglyNumbers) < n {
		checkedInteger := integer

		if checkedInteger%2 == 0 {
			checkedInteger = divideIfDivisible(checkedInteger, 2)
		}
		if checkedInteger%3 == 0 {
			checkedInteger = divideIfDivisible(checkedInteger, 3)
		}
		if checkedInteger%5 == 0 {
			checkedInteger = divideIfDivisible(checkedInteger, 5)
		}

		if checkedInteger == 1 {
			uglyNumbers = append(uglyNumbers, integer)
		}
		integer++
	}

	return uglyNumbers[n-1]

}

func divideIfDivisible(num int, divider int) int {
	if num%divider == 0 {
		num = divideIfDivisible(num/divider, divider)
	}
	return num
}
