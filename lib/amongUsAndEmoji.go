package lib

import (
	"math/rand"
	"time"
)

func AmongUsAndEmoji(size int) string {
	var str string
	for i := 0; i < size; i++ {
		rand.Seed(time.Now().UnixNano() * (int64(i)+1))
		integer := rand.Intn(2)

		switch integer {
		case 0:
			str = str + RandomEmoji(1)
		case 1:
			str = str + AmongUs(1)
		}
	}
	return str
}