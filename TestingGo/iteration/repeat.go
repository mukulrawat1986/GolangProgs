package iteration

// Repeat repeats a character
func Repeat(character string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}

	return repeated
}
