package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	r := 0.0
	nums := []int{}
	pNums1 := len(nums1) - 1
	pNums2 := len(nums2) - 1
	tLength := (pNums1 + 1 + pNums2 + 1)
	lNums := (tLength / 2)
	i := 0
	p1 := 0
	p2 := 0
	for ; i <= lNums; i++ {
		a := 0
		if p1 > pNums1 {
			a = nums2[p2]
			p2++
		} else if p2 > pNums2 {
			a = nums1[p1]
			p1++
		} else if nums1[p1] > nums2[p2] {
			a = nums2[p2]
			p2++
		} else {
			a = nums1[p1]
			p1++
		}
		nums = append(nums, a)
	}
	if tLength%2 == 0 {
		r = float64(nums[lNums-1]+nums[lNums]) / 2
	} else {
		r = float64(nums[lNums])
	}
	return r
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	print(findMedianSortedArrays(nums1, nums2))
}
