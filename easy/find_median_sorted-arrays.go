package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	for i := 0; i < len(nums2); i++ {
		nums1 = append(nums1, nums2[i])
	}

	return 0.0
}
