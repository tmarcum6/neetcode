package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	t := "Was it a car or a cat I saw?"
	b := isPalindrome(t)
	fmt.Println(b)
}

func hasDuplicate(nums []int) bool {
	dupeMap := make(map[int]int)
	for _, v := range nums {
		dupeMap[v]++
		if dupeMap[v] > 1 {
			return true
		}
	}

	return false
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := make(map[rune]int)
	for _, ch := range s {
		sMap[ch]++
	}
	for _, ch := range t {
		sMap[ch]--
	}

	for _, v := range sMap {
		if v != 0 {
			return false
		}
	}

	return true
}

func twoSum(nums []int, target int) []int {
	var s []int
	for i := 0; i < len(nums)-1; i++ {
		pt1 := nums[i]
		for j := i + 1; j < len(nums); j++ {
			pt2 := nums[j]
			if pt1+pt2 == target {
				s = append(s, i)
				s = append(s, j)
				return s
			}
		}
	}

	return s
}

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, s := range strs {
		key := sortStrings(s)
		groups[key] = append(groups[key], s)
	}

	var ret [][]string
	for _, v := range groups {

		ret = append(ret, v)
	}

	return ret
}

func sortStrings(s string) string {
	chars := []rune(s)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	return string(chars)
}

func topKFrequent(nums []int, k int) []int {
	// n = total number of elements in nums
	// m = number of unique elements in nums

	// Time: O(n)
	// We iterate through nums once to build the frequency map.
	// Space: O(m)
	counterMap := make(map[int]int)
	for _, n := range nums {
		counterMap[n]++
	}

	// Time: O(m)
	// We convert the map into a slice of size m.
	// Space: O(m)
	sorted := []KVPair{}
	for key, value := range counterMap {
		sorted = append(sorted, KVPair{key, value})
	}

	// Time: O(m log m)
	// Sorting m unique elements by frequency.
	// Space: O(1) auxiliary (in-place sort)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Value > sorted[j].Value
	})

	// Time: O(k)
	// Extract top k elements.
	// Space: O(k)
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[i] = sorted[i].Key
	}

	// Total Time Complexity:
	// O(n + m log m)
	// Worst case: m = n → O(n log n)
	//
	// Total Space Complexity:
	// O(m + k)
	// Worst case: O(n)

	// Bucket / Heap sort to optimize

	return ret
}

type KVPair struct {
	Key   int
	Value int
}

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	// n = number of strings
	// m = total number of characters across all strings

	// Space: O(m)
	// strings.Builder stores the full encoded output.
	b := strings.Builder{}

	// Time: O(n + m)
	// We iterate through each string once and write:
	// - the string length
	// - separator '#'
	// - the string contents
	for _, s := range strs {
		b.WriteString(strconv.Itoa(len(s))) // O(log len(s))
		b.WriteByte('#')                    // O(1)
		b.WriteString(s)                    // O(len(s))
	}

	// Time: O(m)
	// Converts builder buffer into final string.
	return b.String()
}

func (s *Solution) Decode(encoded string) []string {
	// m = total length of encoded string

	// Space: O(m)
	// Result stores all decoded strings.
	result := []string{}

	// i tracks current parsing position.
	i := 0

	// Time: O(m)
	// Each character in encoded is visited at most once.
	for i < len(encoded) {

		// j scans forward to find '#'
		j := i

		// Time across all iterations: O(m)
		// Finds the separator for current encoded string.
		for j < len(encoded) && encoded[j] != '#' {
			j++
		}

		// Safety check for malformed input.
		if j == len(encoded) {
			break
		}

		// Time: O(length digits)
		// Converts substring length -> integer.
		length, _ := strconv.Atoi(encoded[i:j])

		j++ // skip '#'

		// Bounds safety check.
		if j+length > len(encoded) {
			break
		}

		// Time: O(length)
		// Extract substring and append to result.
		result = append(result, encoded[j:j+length])

		// Move to next encoded segment.
		i = j + length
	}

	// Total Time Complexity:
	// O(m)
	// Every character is processed at most once.
	//
	// Total Space Complexity:
	// O(m)
	// Output storage dominates.

	return result
}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	ret := make([]int, n)

	// n = number of elements in nums

	// --------------------------
	// SPACE COMPLEXITY
	// --------------------------
	// Space: O(1) extra space (excluding output array)
	// ret is required for the output, so it is NOT counted as extra space.
	// Only one extra variable is used: "right" → O(1)

	// --------------------------
	// LEFT PASS (prefix products)
	// --------------------------
	// Time: O(n)
	// Space: O(1) extra
	// We store prefix products directly in ret.
	ret[0] = 1
	for i := 1; i < n; i++ {
		ret[i] = ret[i-1] * nums[i-1]
	}

	// --------------------------
	// RIGHT PASS (suffix products)
	// --------------------------
	// Time: O(n)
	// Space: O(1) extra
	// We use a single variable "right" to accumulate suffix product.
	right := 1
	for i := n - 1; i >= 0; i-- {
		ret[i] *= right
		right *= nums[i]

		// Debug only
		// fmt.Println(right)
	}

	// --------------------------
	// TOTAL COMPLEXITY
	// --------------------------
	// Time: O(n)
	// Space: O(1) extra space (excluding output array)
	// Reason: we only use constant extra variables regardless of input size

	return ret
}

func longestConsecutive(nums []int) int {
	//must be O(n) - one loop

	//n = number of integers in nums
	//m = space for each index in nums

	//can be negative

	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}

	longest := 0
	for v := range m {
		//if the number before v does not exist we are at the start of a seq
		if !m[v-1] {
			current := v
			streak := 1

			//continue until seq chain is broken
			for m[current+1] {
				current++
				streak++
			}

			if streak > longest {
				longest = streak
			}
		}
	}

	return longest
}

func isPalindrome(s string) bool {
	b := strings.Builder{}

	for i := 0; i < len(s); i++ {
		c := s[i]

		if c >= 'a' && c <= 'z' ||
			c >= 'A' && c <= 'Z' ||
			c >= '0' && c <= '9' {
			b.WriteByte(c)
		}
	}

	filteredString := b.String()
	filteredString = strings.ToLower(filteredString)

	i := 0
	j := len(filteredString) - 1
	r := []rune(filteredString)

	for i < j {
		if r[i] != r[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func threeSum(nums []int) [][]int {

	return [][]int{}
}
