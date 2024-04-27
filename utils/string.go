package utils

// Reverses a given string.
// converts them to rune first for even encoding.
// from : https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
