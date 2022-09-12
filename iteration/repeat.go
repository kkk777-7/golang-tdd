package iteration

const repeatCount = 5

func Repeat(word string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += word
	}
	return repeated
}
