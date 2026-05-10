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

// func TestTopKFrequent(t *testing.T) {
// 	nums := []int{1, 2, 2, 3, 3, 3}
// 	k := 2
// 	expected := []int{2, 3}
// 	result := topKFrequent(nums, k)
//
// 	if !slices.Equal(result, expected) {
// 		t.Errorf("topKFrequent= %v; want %v", result, k)
// 	}
// }

func TestProductExceptSelf(t *testing.T) {
	nums := []int{1, 2, 4, 6}
	expected := []int{48, 24, 12, 8}
	result := productExceptSelf(nums)

	if !slices.Equal(result, expected) {
		t.Errorf("productExceptSelf= %v; want %v", result, expected)
	}
}

// func TestLongestConsecutive(t *testing.T) {
// 	nums := []int{2, 20, 4, 10, 3, 4, 5}
// 	expected := 5
// 	result := longestConsecutive(nums)
//
// 	if expected != result {
// 		t.Errorf("longestConsecutive= %v; want %v", result, expected)
// 	}
// }

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	expected := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	result := threeSum(nums)

	for i := range result {
		if !slices.Equal(result[i], expected[i]) {
			t.Errorf("threeSum= %v; want %v", result, expected)
		}
	}
}

func TestIsValidSudoku(t *testing.T) {
	board1 := [][]byte{
		{'1', '2', '.', '.', '3', '.', '.', '.', '.'},
		{'4', '.', '.', '5', '.', '.', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '.', '3'},
		{'5', '.', '.', '.', '6', '.', '.', '.', '4'},
		{'.', '.', '.', '8', '.', '3', '.', '.', '5'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '.', '.', '.', '.', '.', '2', '.', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '8'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	expected1 := true
	result1 := isValidSudoku(board1)

	if result1 != expected1 {
		t.Errorf("isValidSudoku=%v; want %v", result1, expected1)
	}

	board2 := [][]byte{
		{'1', '2', '.', '.', '3', '.', '.', '.', '.'},
		{'4', '.', '.', '5', '.', '.', '.', '.', '.'},
		{'.', '9', '1', '.', '.', '.', '.', '.', '3'},
		{'5', '.', '.', '.', '6', '.', '.', '.', '4'},
		{'.', '.', '.', '8', '.', '3', '.', '.', '5'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '.', '.', '.', '.', '.', '2', '.', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '8'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	expected2 := false
	result2 := isValidSudoku(board2)

	if result1 != expected1 {
		t.Errorf("isValidSudoku=%v; want %v", result2, expected2)
	}
}

func TestCheckMagazine(t *testing.T) {
	magazine := []string{"give", "me", "one", "grand", "today", "night"}
	note := []string{"give", "one", "grand", "today"}
	expected := true
	result := checkMagazine(note, magazine)

	if result != expected {
		t.Errorf("checkMagazine=%v; want %v", result, expected)
	}

	magazine2 := []string{"two", "times", "three", "is", "not", "four"}
	note2 := []string{"two", "times", "two", "is", "four"}
	result2 := checkMagazine(note2, magazine2)
	expected2 := false

	if result2 != expected2 {
		t.Errorf("checkMagazine=%v; want %v", result2, expected2)
	}
}
