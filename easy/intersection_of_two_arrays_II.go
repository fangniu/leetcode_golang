package main

//Given two arrays, write a function to compute their intersection.
//
//Example:
//Given nums1 = [1, 2, 2, 1], nums2 = [2, 2], return [2].
//
//Note:
//Each element in the result must be unique.
//The result can be in any order.


func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}
	m := map[int]int{}
	var ret []int
	for _, e:= range nums1 {
		if v, ok := m[e]; ok {
			m[e] = v + 1
		} else {
			m[e] = 1
		}
	}
	for _, e:= range nums2 {
		if v, ok := m[e]; ok {
			ret = append(ret, e)
			v --
			if v == 0 {
				delete(m, e)
			} else {
				m[e] = v
			}
		}
	}
	return ret
}