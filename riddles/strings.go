package riddles

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

// https://www.geeksforgeeks.org/string-data-structure/

// sumAllNumbers
// Given a string containing alphanumeric characters,
// calculate sum of all numbers present in the string.
func sumAllNumbers(str string) int {
	if len(str) == 0 {
		return 0
	}

	if unicode.IsDigit(rune(str[0])) {
		value, _ := strconv.Atoi(string(str[0]))
		return value + sumAllNumbers(str[1:])
	}

	return sumAllNumbers(str[1:])
}

// maxNumeric
// Given an alphanumeric string,
// extract maximum numeric value from that string.
func maxNumeric(str string) int {
	if len(str) == 0 {
		return 0
	}

	if unicode.IsDigit(rune(str[0])) {
		value, _ := strconv.Atoi(string(str[0]))
		max := math.Max(float64(value), float64(maxNumeric(str[1:])))
		return int(max)
	}
	return maxNumeric(str[1:])
}

// reverseInPlace
// duplicate characters from a string
func duplicatedChars(str string) []string {
	dups := make(map[rune]int)
	for _, v := range str {
		dups[v] = dups[v] + 1
	}

	var res []string
	for k, v := range dups {
		if v > 1 {
			res = append(res, string(k))
		}
	}
	return res
}

// areAnagrams
// two strings are anagrams of each other
func areAnagrams(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	acc := make(map[uint8]int)
	for i := 0; i < len(str1); i++ {
		acc[str1[i]] = acc[str1[i]] + 1
		acc[str2[i]] = acc[str2[i]] + 1
	}

	for _, v := range acc {
		if math.Mod(float64(v), 2) != 0 {
			return false
		}
	}
	return true
}

// allPermutations
// find all the permutations of a string
func allPermutations(str string, left int) []string {
	if left == len(str)-1 {
		return []string{str}
	}

	var strs []string
	for i := left; i < len(str); i++ {
		runes := []rune(str)
		runes[i], runes[left] = runes[left], runes[i]

		newStr := string(runes)
		x := allPermutations(newStr, left+1)
		strs = append(strs, x...)
	}
	return strs
}

// reverseRecursion
// given string be reversed using recursion
func reverseRecursion(str string) string {
	if len(str) <= 1 {
		return str
	}

	return str[len(str)-1:] + reverseRecursion(str[:len(str)-1])
}

// count the occurrence of a given character in a string
func countCharOccurrences(str string, char uint8) int {
	if len(str) == 0 {
		return 0
	}

	if str[0] == char {
		return 1 + countCharOccurrences(str[1:], char)
	}
	return countCharOccurrences(str[1:], char)
}

// firstNonRepeatable
// first non-repeated character from a string
func firstNonRepeatable(str string) int32 {
	m := make(map[int32]int)
	for _, v := range str {
		m[v] = m[v] + 1
	}

	for _, v := range str {
		if m[v] == 1 {
			return v
		}
	}
	return 0
}

// reverseWords
// reverse words in a given sentence
func reverseWords(str string) string {
	split := strings.Split(str, " ")
	for i := 0; i < len(split)/2; i++ {
		split[i], split[len(split)-1-i] = split[len(split)-1-i], split[i]
	}
	return strings.Join(split, " ")
}

// areRotation
// check if two strings are a rotation of each other
func areRotation(strA, strB string) bool {
	if len(strA) != len(strB) {
		return false
	}

	rB := []rune(strB)
	for i := 0; i < len(strA); i++ {
		if strA == string(rB) {
			return true
		}
		rB = append(rB[len(strA)-1:], rB[:len(strA)-1]...)
	}
	return false
}

// isPalindrome
// check if a given string is a palindrome
func isPalindrome(str string) bool {
	if len(str) <= 1 {
		return true
	}

	if str[0] == str[len(str)-1] {
		return isPalindrome(str[1 : len(str)-1])
	}

	return false
}

// longestPalindrome
// find the longest palindromic substring in str
func longestPalindrome(str string) string {
	if isPalindrome(str) {
		return str
	}

	maxLength := 1
	start := 0
	for i := 0; i < len(str); i++ {
		for j := i; j < len(str); j++ {
			flag := 1

			// Check palindrome
			for k := 0; k < (j-i+1)/2; k++ {
				if str[i+k] != str[j-k] {
					flag = 0
				}
			}

			// Palindrome
			if flag > 0 && (j-i+1) > maxLength {
				start = i
				maxLength = j - i + 1
			}
		}
	}
	fmt.Println(str, start, start+maxLength-1)
	return str[start : start+maxLength]
}

// findLongestSubstring
// find the length of the longest substring without repeating characters
func findLongestSubstring(str string) string {
	maxStrs := []string{""}
	var i int
	for _, c := range str {
		if strings.ContainsRune(maxStrs[i], c) {
			i++
			maxStrs = append(maxStrs, "")
		}
		maxStrs[i] += string(c)
	}

	for j := 0; j < len(maxStrs)-1; j++ {
		if len(maxStrs[j]) > len(maxStrs[j+1]) {
			maxStrs[j], maxStrs[j+1] = maxStrs[j+1], maxStrs[j]
		}
	}

	return maxStrs[len(maxStrs)-1]
}

// removeDuplicate
// remove the duplicate character from String
func removeDuplicate(str string) string {
	if len(str) <= 0 {
		return str
	}
	if strings.Contains(str[1:], string(str[0])) {
		return removeDuplicate(str[1:])
	}
	return string(str[0]) + removeDuplicate(str[1:])
}

// findMaxWordOccurring
// Given an array of strings, find the most frequent word in a given array
func findMaxWordOccurring(text []string) string {
	counter := make(map[string]int)
	for _, str := range text {
		for _, word := range strings.Split(str, " ") {
			counter[word] += 1
		}
	}

	var word string
	var count int
	for k, v := range counter {
		if v > count {
			count = v
			word = k
		}
	}
	return word
}
