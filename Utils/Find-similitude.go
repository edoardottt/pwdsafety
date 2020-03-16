package Utils


func FindExactly(words []string, password string) bool {
	for _, s := range words {
		if password==s+"\n" {
			return true
		}
	}
	return false
}