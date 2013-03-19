package bitmask

/* Return list of the flags set in the bitmask
*/
func FlagsSet(f int) (setFlags []int) {
	c := 1
	for c <= f {
		if c & f == c {
			setFlags = append(setFlags, c)
		}
		c = c * 2
	}
	return
}
