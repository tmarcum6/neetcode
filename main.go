package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"act", "pots", "tops", "cat", "stop", "hat"}
	ret := groupAnagrams(strs)
	fmt.Println(ret)
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
	//sorts if i is < j (alphabetically)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	return string(chars)
}
