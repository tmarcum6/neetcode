package main

import (
	"slices"
	"sort"
	"testing"
)

func TestHasDuplicate(t *testing.T) {
	nums := []int{1, 2, 3, 3}
	result := hasDuplicate(nums)
	expected := true

	if result != expected {
		t.Errorf("hasDuplicate = %v; want %v", result, expected)
	}
}

func TestIsAnagram(t *testing.T) {
	s1 := "x"
	s2 := "xx"
	result := isAnagram(s1, s2)
	expected := false

	if result != expected {
		t.Errorf("isAnagram = %v; want %v", result, expected)
	}
}

func TestTwoSum(t *testing.T) {
	nums := []int{2, 5, 5, 11}
	target := 10
	result := twoSum(nums, target)
	expected := []int{1, 2}

	if !slices.Equal(result, expected) {
		t.Errorf("twoSum = %v; want %v", result, expected)
	}
}

func TestGroupAnagrams(t *testing.T) {
	strs := []string{"act", "pots", "tops", "cat", "stop", "hat"}
	result := groupAnagrams(strs)
	expected := [][]string{{"hat"}, {"act", "cat"}, {"stop", "pots", "tops"}}

	// normalize both groups
	sortGroup(result)
	sortGroup(expected)

	if len(result) != len(expected) {
		t.Errorf("groupAnagrams = %v; want %v", result, expected)
	}

	for i := range result {
		//final sorting of each string
		sort.Strings(result[i])
		sort.Strings(expected[i])

		if len(result[i]) != len(expected[i]) {
			t.Fatalf("got %v, want %v", result, expected)
		}

		for j := range result[i] {
			if result[i][j] != expected[i][j] {
				t.Fatalf("got %v, want %v", result, expected)
			}
		}
	}
}

func sortGroup(s [][]string) {
	sort.Slice(s, func(i, j int) bool {
		return s[i][0] < s[j][0]
	})
}
