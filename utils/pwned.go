package utils

import (
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// Result describes a result from the Pwned Password service.
type Result struct {
	// Pwned has the password been seen at least once. A value of false doesn't mean the password is any good though.
	Pwned bool
	// TimesObserved the number of times this password has been seen by the pwned password service.
	TimesObserved uint64
}

type pwnedHash struct {
	Hash  string
	Range string
}

// IsPwned will synchronously check if the provided password has been pwned.
func IsPwned(password string) (*Result, error) {
	if password == "" {
		return nil, fmt.Errorf("empty password provided")
	}

	hash, err := getHash(password)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get("https://api.pwnedpasswords.com/range/" + hash.Range)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(body), "\r\n")
	for _, line := range lines {
		components := strings.Split(line, ":")
		if len(components) != 2 {
			return nil, fmt.Errorf("invalid response from pwned password API")
		}

		resultHash := components[0]
		countStr := components[1]

		if hash.Range+resultHash == hash.Hash {
			count, err := strconv.ParseUint(countStr, 10, 64)
			if err != nil {
				return nil, err
			}

			ret := Result{
				Pwned:         true,
				TimesObserved: count,
			}
			return &ret, nil
		}
	}

	ret := Result{
		Pwned:         false,
		TimesObserved: 0,
	}
	return &ret, nil
}

func getHash(password string) (*pwnedHash, error) {
	h := sha1.New()
	_, err := io.WriteString(h, password)
	if err != nil {
		return nil, err
	}
	hash := fmt.Sprintf("%x", h.Sum(nil))
	hash = strings.ToUpper(hash)
	if len(hash) < 5 {
		return nil, fmt.Errorf("unable to hash password")
	}

	result := pwnedHash{
		Hash:  hash,
		Range: hash[0:5],
	}

	return &result, nil
}

//Reverse : Reverse the input string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
