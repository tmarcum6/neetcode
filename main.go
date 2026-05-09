package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := []string{"Hello", "World"}
	s := Solution{}
	t := s.Encode(input)
	k := s.Decode(t)
	fmt.Println(k)

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
	b := strings.Builder{}

	for _, s := range strs {
		b.WriteString(strconv.Itoa(len(s)))
		b.WriteByte('#')
		b.WriteString(s)
	}

	return b.String()
}

func (s *Solution) Decode(encoded string) []string {
	result := []string{}
	i := 0
	for i < len(encoded) {
		j := i
		for j < len(encoded) && encoded[j] != '#' {
			j++
		}

		if j == len(encoded) {
			break
		}

		length, _ := strconv.Atoi(encoded[i:j])

		j++

		if j+length > len(encoded) {
			break
		}

		result = append(result, encoded[j:j+length])

		i = j + length
	}

	return result
}
