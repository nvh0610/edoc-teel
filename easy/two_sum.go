package main

//func main() {
//	nums := []int{2, 7, 11, 15}
//	target := 9
//	fmt.Println(twoSum(nums, target))
//}

func TwoSum(nums []int, target int) []int {
	m := map[int]int{}

	for i, num := range nums {
		if key, ok := m[target-num]; ok {
			return []int{i, key}
		}

		m[num] = i
	}

	return []int{}
}
