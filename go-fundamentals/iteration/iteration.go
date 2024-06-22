package iteration

const repeatedCount = 5

func Repeat(toRepeat string) string {
	repeatedOut := ""
	for i := 0; i < repeatedCount; i++ {
		repeatedOut += toRepeat
	}
	return repeatedOut
}
