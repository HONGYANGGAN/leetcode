func minWindow(s string, t string) string {
	m, n := len(s), len(t)
	if n > m {
		return ""
	}
	need, window := make(map[byte]int), make(map[byte]int)
	needCount, windowCount := 0, 0
	for _, r := range t {
		if need[byte(r)] == 0 {
			needCount++
		}
		need[byte(r)]++
	}
	start, minLen := 0, math.MaxInt32
	left, right := 0, 0
	for right < m {
		ch := s[right]
		right++
		if v, ok := need[ch]; ok {
			window[ch]++
			if window[ch] == v {
				windowCount++
			}
		}
		for windowCount == needCount {
			if right-left < minLen {
				minLen = right - left
				start = left
			}
			ch = s[left]
			left++
			if v, ok := need[ch]; ok {
				if window[ch] == v {
					windowCount--
				}
				window[ch]--
			}
		}
	}
	if minLen == math.MaxInt32 {
		return ""
	}
	return s[start : start+minLen]
}