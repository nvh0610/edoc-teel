package main

func LengthOfLongestSubstring(s string) int {
	m := map[uint8]int{}
	i, j, count, maxCount := 0, 0, 0, 0
	for i < len(s) {
		str := s[i]
		if _, ok := m[str]; !ok {
			count += 1
			if count > maxCount {
				maxCount = count
			}
			i += 1
			m[str] = i
		} else {
			j += m[str]
			i = j
			count = 0
			for key, _ := range m {
				delete(m, key)
			}
		}
	}
	return maxCount
}
