package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	end := n + m - 1
	m--
	n--
	for n >= 0 {
		if m >= 0 && nums1[m] >= nums2[n] {
			nums1[end] = nums1[m]
			m--
		} else {
			nums1[end] = nums2[n]
			n--
		}
		end--
	}
}
