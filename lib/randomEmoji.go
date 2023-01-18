package lib

import (
	"html"
	"math"
	"math/rand"
	"strconv"
	"time"
)

// (c) 2020 Vladimir Babin (modify Kroks in order to support a determined length of emojis :joy: :joy:)
// This code is licensed under MIT license.

func RandomEmoji(size int) string {
	emoji := [][]int{
		{128513, 128591},
		{128640, 128704},
	}
	var str string
	for i := 0; i < size; i++ {
		rand.Seed(time.Now().UnixNano() * (int64(i) + 1))
		r := emoji[rand.Int()%len(emoji)]
		min := r[0]
		max := r[1]
		n := rand.Intn(max-min+1) + min
		str = str + html.UnescapeString("&#"+strconv.Itoa(n)+";") + InvisibleUrl(int(math.Round(float64(16/size))))
	}
	return str
}
