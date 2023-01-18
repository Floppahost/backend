package lib

import (
	"html"
	"math/rand"
	"strconv"
	"time"
)

func InvisibleUrl(size int) string {
	var chars [3]int
	chars[0] = 8203
	chars[1] = 8204
	chars[2] = 8205

	var str string
	for i := 0; i < size; i++ {
		rand.Seed(time.Now().UnixNano() * (int64(i) + 1))
		integer := rand.Intn(3)
		char := chars[integer]
		str = str + html.UnescapeString("&#"+strconv.Itoa(char)+";")
	}
	return str
}
