package cmp

// CompareOnStringSlice compare if two string slice is equal
func CompareOnStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	m := make(map[string]int, len(b))
	for _, v := range b {
		m[v]++
	}

	isEqual := true
	for _, v := range a {
		if _, ok := m[v]; ok {
			m[v]--
			if m[v] <= 0 {
				delete(m, v)
			}
		} else {
			isEqual = false
			break
		}
	}

	return isEqual && len(m) == 0
}
