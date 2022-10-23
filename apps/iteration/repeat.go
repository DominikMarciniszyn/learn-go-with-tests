package iteration

func Repeat(character string, repeatCount int) string {
	var repeated string

	for index := 0; index < repeatCount; index++ {
		repeated += character
	}

	return repeated
}
