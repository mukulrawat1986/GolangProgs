package iteration

const repeatCount = 5

// Repeat function repeates a string five times
func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
