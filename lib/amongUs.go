package lib

import "math"

func AmongUs(size int) string {
	char := "ඞ"
	var str string

	for i := 0; i < size; i++ {
		str = str + char + InvisibleUrl(int(math.Round(float64(25/size))))
	}
	return str
}
