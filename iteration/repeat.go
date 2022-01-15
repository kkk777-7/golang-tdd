package iteration

func Repeat(word string) (res string) {
	for i := 0; i < 5; i++ {
		res += word
	}
	return res
}
