package iteration

const repeatConst = 5

func Repeat(char string) string {
	var repeated string

	for i := 0; i < repeatConst; i++ {
		repeated += char
	}

	return repeated
}
