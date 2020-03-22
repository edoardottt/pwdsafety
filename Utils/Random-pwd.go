package Utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func GenerateRandom(engWords []string) string {
	rand.Seed(time.Now().UnixNano())
	max := len(engWords) - 1
	rand1 := rand.Intn(max)
	rand2 := rand.Intn(max)
	word1 := engWords[rand1]
	word2 := engWords[rand2]
	rand4 := rand.Intn(100)
	newPwd := word1 + "-" + strings.ToUpper(word2) + "-" + strconv.Itoa(rand4)
	return newPwd
}
