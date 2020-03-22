package Utils

//Check if the password is equal to one known password
func FindExactly(words []string, password string) bool {
	for _, s := range words {
		if password == s {
			return true
		}
	}
	return false
}

//Reverse the input string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

//Check if the password is equal to one known password reversed
func FindExactlyReversed(words []string, password string) bool {
	for _, s := range words {
		if Reverse(password) == s {
			return true
		}
	}
	return false
}
