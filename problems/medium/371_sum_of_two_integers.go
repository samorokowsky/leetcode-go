package medium

import (
	"fmt"
	"github.com/samorokowsky/leetcode-go/util"
	"math"
	"strconv"
	"strings"
)

// Problem371 (https://leetcode.com/problems/sum-of-two-integers/description/)
/*
Given two integers a and b, return the sum of the two integers without using the operators + and -.

Example 1:
Input: a = 1, b = 2
Output: 3

Example 2:
Input: a = 2, b = 3
Output: 5

Constraints:
-1000 <= a, b <= 1000
*/
func Problem371() {
	var examples = map[util.Pair]int{
		util.Pair{First: 1, Second: 2}:      3,
		util.Pair{First: 2, Second: 3}:      5,
		util.Pair{First: 1000, Second: 999}: 1999,
		util.Pair{First: 3, Second: -8}:     -5,
	}

	for k, v := range examples {
		result := getSum(k.First, k.Second)
		if result != v {
			var log = fmt.Sprintf("Problem371.getSum - result from '%d' = %d doesn't match expected %d.",
				k, result, v)
			panic(log)
		}
	}

	fmt.Println(" Problem371 - Success passed")

}

func getSum(a int, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	binaryAStr := strings.ReplaceAll(fmt.Sprintf("%16b", uint16(a)), " ", "")
	binaryBStr := strings.ReplaceAll(fmt.Sprintf("%16b", uint16(b)), " ", "")

	resultBinaryString := sumTwoBinaryStrings(binaryAStr, binaryBStr, a > 0 && b > 0)

	number, _ := strconv.ParseUint(resultBinaryString, 2, 16)

	if number >= 0x8000 {
		number = number - 0x10000
	}

	return int(number)
}

func sumTwoBinaryStrings(binaryAStr string, binaryBStr string, allPositive bool) string {
	binaryAStr, binaryBStr = getEqualBinaryStringsByLength(binaryAStr, binaryBStr)

	transfer := 0
	var resultBinarySlice []int

	for i := len(binaryAStr) - 1; i >= 0; i-- {
		aElementAtIndex := int(binaryAStr[i] - '0')
		bElementAtIndex := int(binaryBStr[i] - '0')
		if aElementAtIndex == 1 && bElementAtIndex == 1 {
			if transfer == 0 {
				resultBinarySlice = append(resultBinarySlice, 0)
				transfer = 1
			} else {
				resultBinarySlice = append(resultBinarySlice, 1)
				transfer = 1
			}
		} else if aElementAtIndex == 0 && bElementAtIndex == 0 {
			if transfer == 0 {
				resultBinarySlice = append(resultBinarySlice, 0)
			} else {
				resultBinarySlice = append(resultBinarySlice, 1)
				transfer = 0
			}
		} else if (aElementAtIndex == 0 && bElementAtIndex == 1) || (aElementAtIndex == 1 && bElementAtIndex == 0) {
			if transfer == 1 {
				resultBinarySlice = append(resultBinarySlice, 0)
			} else {
				resultBinarySlice = append(resultBinarySlice, 1)
			}
		}
	}

	if transfer == 1 && allPositive {
		resultBinarySlice = append(resultBinarySlice, 1)
	}

	return transformBinarySliceToBinaryString(reverseSlice(resultBinarySlice))
}

func transformBinarySliceToBinaryString(binarySlice []int) string {
	var strSlice []string
	for _, num := range binarySlice {
		strSlice = append(strSlice, strconv.Itoa(num))
	}

	return strings.Join(strSlice, "")
}

func getEqualBinaryStringsByLength(binaryAStr string, binaryBStr string) (string, string) {

	byteArrayA := []byte(binaryAStr)
	byteArrayB := []byte(binaryBStr)

	if len(byteArrayA) > len(byteArrayB) {
		for len(byteArrayA) > len(byteArrayB) {
			byteArrayB = append([]byte{'0'}, byteArrayB...)
		}
	}
	if len(byteArrayB) > len(byteArrayA) {
		for len(byteArrayB) > len(byteArrayA) {
			byteArrayA = append([]byte{'0'}, byteArrayA...)
		}
	}

	return string(byteArrayA), string(byteArrayB)
}

func reverseSlice(slice []int) []int {
	for i := 0; i < len(slice)/2; i++ {
		slice[i], slice[len(slice)-1-i] = slice[len(slice)-1-i], slice[i]
	}
	return slice
}

func getSumWrongVersion(a int, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	expPowA := math.Pow(2, float64(a))
	expPowB := math.Pow(2, float64(b))

	return int(math.Log(expPowA*expPowB) / math.Log(2))
}
