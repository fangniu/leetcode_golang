package main

import "fmt"

//Suppose Andy and Doris want to choose a restaurant for dinner, and they both have a list of favorite restaurants represented by strings.
//
//You need to help them find out their common interest with the least list index sum. If there is a choice tie between answers, output all of them with no order requirement. You could assume there always exists an answer.
//
//Example 1:
//Input:
//["Shogun", "Tapioca Express", "Burger King", "KFC"]
//["Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"]
//Output: ["Shogun"]
//Explanation: The only restaurant they both like is "Shogun".
//Example 2:
//Input:
//["Shogun", "Tapioca Express", "Burger King", "KFC"]
//["KFC", "Shogun", "Burger King"]
//Output: ["Shogun"]
//Explanation: The restaurant they both like and have the least index sum is "Shogun" with index sum 1 (0+1).
//Note:
//The length of both lists will be in the range of [1, 1000].
//The length of strings in both lists will be in the range of [1, 30].
//The index is starting from 0 to the list length minus 1.
//No duplicates in both lists.

func findRestaurant(list1 []string, list2 []string) []string {
	var res []string
	min := len(list2) + len(list1)
	m := map[string]int{}
	for i, s := range list1 {
		m[s] = i
	}
	for i, s := range list2 {
		if v, ok := m[s]; ok {
			if v + i < min {
				min = v + i
				res = res[:0]
				res = append(res, s)
			} else if v + i == min {
				res = append(res, s)
			}
		} else {
			m[s] = i
		}
	}
	return res
}

func main() {
	fmt.Println(findRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
	[]string{"KFC", "Shogun", "Burger King"}))
}