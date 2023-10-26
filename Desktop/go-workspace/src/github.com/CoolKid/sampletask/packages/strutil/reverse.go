package strutil

// Reverse returns its argument string reversed rune-wise left to right.

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	runes := []rune(s)
	for i,j := 0, len(runes)-1; i < len(runes)/2; i,j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}