package iteration

func Repeat(toRepeat string, amount int) string {
	var repeatedOut string
	for i := 0; i < amount; i++ {
		repeatedOut += toRepeat
	}
	return repeatedOut

	// return strings.Repeat(toRepeat, amount)
}
