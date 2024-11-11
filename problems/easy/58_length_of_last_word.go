package easy

import "fmt"

// Problem58 (https://leetcode.com/problems/length-of-last-word/description/)
/*
Given a string s consisting of words and spaces, return the length of the last word in the string.

A word is a maximal
substring
consisting of non-space characters only.

Example 1:
Input: s = "Hello World"
Output: 5
Explanation: The last word is "World" with length 5.

Example 2:
Input: s = "   fly me   to   the moon  "
Output: 4
Explanation: The last word is "moon" with length 4.

Example 3:
Input: s = "luffy is still joyboy"
Output: 6
Explanation: The last word is "joyboy" with length 6.

Constraints:

1 <= s.length <= 104
s consists of only English letters and spaces ' '.
There will be at least one word in s.
*/
func Problem58() {
	var examples = map[string]int{
		"Hello World":                 5,
		"   fly me   to   the moon  ": 4,
		"luffy is still joyboy":       6,
	}

	for k, v := range examples {
		result := lengthOfLastWord(k)
		if result != v {
			var log = fmt.Sprintf("Problem58.lengthOfLastWord - result from '%s' = %d doesn't match expected %d.",
				k, result, v)
			panic(log)
		}
	}

	fmt.Println("Problem58 - Success passed")

}

func lengthOfLastWord(s string) int {
	runes := []rune(s)
	isFirstSpaces := true
	letterCount := 0
	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] != ' ' && isFirstSpaces {
			isFirstSpaces = false
			letterCount++
		} else if runes[i] == ' ' && isFirstSpaces {
			continue
		} else if runes[i] != ' ' && !isFirstSpaces {
			letterCount++
		} else if runes[i] == ' ' && !isFirstSpaces {
			break
		}
	}

	return letterCount
}
