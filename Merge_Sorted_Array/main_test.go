package main

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	expected := []int{1, 2, 2, 3, 5, 6}
	merge(nums1, m, nums2, n)
	if !reflect.DeepEqual(nums1, expected) {
		t.Error("\nexpected:", expected, "\ngot", nums1)
	}
}

func Test2(t *testing.T) {
	nums1 := []int{0, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	expected := []int{0, 2, 2, 3, 5, 6}
	merge(nums1, m, nums2, n)
	if !reflect.DeepEqual(nums1, expected) {
		t.Error("\nexpected:", expected, "\ngot", nums1)
	}
}

func Test3(t *testing.T) {
	nums1 := []int{1, 0}
	m := 1
	nums2 := []int{0}
	n := 1
	expected := []int{0, 1}
	merge(nums1, m, nums2, n)
	if !reflect.DeepEqual(nums1, expected) {
		t.Error("\nexpected:", expected, "\ngot", nums1)
	}
}

func Test4(t *testing.T) {
	nums1 := []int{0, 0}
	m := 1
	nums2 := []int{1}
	n := 1
	expected := []int{0, 1}
	merge(nums1, m, nums2, n)
	if !reflect.DeepEqual(nums1, expected) {
		t.Error("\nexpected:", expected, "\ngot", nums1)
	}
}

func Test5(t *testing.T) {
	nums1 := []int{0, 0}
	m := 1
	var nums2 []int
	n := 0
	expected := []int{0, 0}
	merge(nums1, m, nums2, n)
	if !reflect.DeepEqual(nums1, expected) {
		t.Error("\nexpected:", expected, "\ngot", nums1)
	}
}
