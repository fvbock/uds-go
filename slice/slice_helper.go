package slice

// TODO: add a version that uses reflection for the lazy ones

// func delete(s []interface{}, i int) {
// 	if len(s) > 2 {
// 		s = append(s[:i], s[i+1:]...)
// 	} else if len(s) == 2 {
// 		s = s[1-i]
// 	} else {
// 		s = []string{}
// 	}
// }

func DeleteIntSlice(s []int, i int) []int {
	// dont want to have a multi val return. we ignore invalid indexes
	if i < 0 || i >= len(s) {
		return s
	}
	if len(s) > 2 {
		s = append(s[:i], s[i+1:]...)
	} else if len(s) == 2 {
		s = []int{s[1-i]}
	} else {
		s = []int{}
	}

	return s
}

func DeleteUint32Slice(s []uint32, i int) []uint32 {
	// dont want to have a multi val return. we ignore invalid indexes
	if i < 0 || i >= len(s) {
		return s
	}
	if len(s) > 2 {
		s = append(s[:i], s[i+1:]...)
	} else if len(s) == 2 {
		s = []uint32{s[1-i]}
	} else {
		s = []uint32{}
	}

	return s
}

func DeleteFloatSlice(s []float64, i int) []float64 {
	// dont want to have a multi val return. we ignore invalid indexes
	if i < 0 || i > len(s) {
		return s
	}
	if len(s) > 2 {
		s = append(s[:i], s[i+1:]...)
	} else if len(s) == 2 {
		s = []float64{s[1-i]}
	} else {
		s = []float64{}
	}

	return s
}

func DeleteStringSlice(s []string, i int) []string {
	// dont want to have a multi val return. we ignore invalid indexes
	if i < 0 || i > len(s) {
		return s
	}
	if len(s) > 2 {
		s = append(s[:i], s[i+1:]...)
	} else if len(s) == 2 {
		s = []string{s[1-i]}
	} else {
		s = []string{}
	}

	return s
}
