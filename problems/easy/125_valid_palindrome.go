package easy

import (
	"fmt"
	"regexp"
	"strings"
)

// Problem125 (https://leetcode.com/problems/valid-palindrome/description/)
/*
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters
and removing all non-alphanumeric characters, it reads the same forward and backward.
Alphanumeric characters include letters and numbers.
Given a string s, return true if it is a palindrome, or false otherwise.

Example 1:
Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.

Example 2:
Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.

Example 3:
Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.

Constraints:
1 <= s.length <= 2 * 105
s consists only of printable ASCII characters.
*/
func Problem125() {
	var examples = map[string]bool{
		"A man, a plan, a canal: Panama": true,
		"race a car":                     false,
		" ":                              true,
	}

	for k, v := range examples {
		result := isPalindrome(k)
		if result != v {
			var log = fmt.Sprintf("Problem125.isPalindrome - result from '%d' = %d doesn't match expected %d.",
				k, result, v)
			panic(log)
		}
	}

	fmt.Println(" Problem125 - Success passed")

}

func isPalindrome(s string) bool {
	alphaNumeric := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	s = strings.ToLower(alphaNumeric.ReplaceAllString(s, ""))
	result := true

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			result = false
			break
		}
	}
	return result
}
