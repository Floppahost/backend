package lib

func AmongUs(size int) string {
	char := "ඞ"
	var str string

	for i := 0; i < size; i++ {
		str = str + char + InvisibleUrl(1)
	}
	return str
}