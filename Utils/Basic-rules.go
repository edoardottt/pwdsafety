package Utils

//Generates a set of unique characters in the input string
func GenerateSetString(input string) []byte {
	set := map[byte]bool{}
	result := []byte{}
	for i := 0; i < len(input); i++ {
		char := input[i]
		existence := set[char]
		if !existence {
			set[char] = true // add element
			result = append(result, char)
		}
	}
	return result
}

//Returns the number of differents characters used in password
func HowManyDifferents(password string) int {
	return len(GenerateSetString(password))
}

//Checks if there is at least one UPPERCASE character
func IsThereUpperCase(password string) bool {
	for _, r := range password {
		if r >= 'A' && r <= 'Z' {
			return true
		}
	}
	return false
}

//Checks if there is at least one lowercase character
func IsThereLowerCase(password string) bool {
	for _, r := range password {
		if r >= 'a' && r <= 'z' {
			return true
		}
	}
	return false
}

//Checks if there is at least one symbol
func IsThereSymbol(password string) bool {
	for _, r := range password {
		if (r < 'A' || r > 'z') && (r < '0' || r > '9' ) {
			return true
		}
	}
	return false
}

//Checks if there is at least one number
func IsThereNumber(password string) bool {
	for _, r := range password {
		if r >= '0' && r <= '9' {
			return true
		}
	}
	return false
}

//Returns how many different types there are in the password
func HowManyTypes(password string) int {
	var howMany int
	if IsThereNumber(password) {howMany++}
	if IsThereUpperCase(password) {howMany++}
	if IsThereLowerCase(password) {howMany++}
	if IsThereSymbol(password) {howMany++}
	return howMany
}