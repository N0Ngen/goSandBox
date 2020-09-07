package iteration

// Repeat tackes one character and returned it five times
func Repeat(character string, times int) string {
	var repeated string
	for i := 0; i < times; i++ {
		repeated += character
	}
	return repeated
}
