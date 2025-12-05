package piscine

func Compact(ptr *[]string) int {
	if ptr == nil {
		return 0
	}
	slice := *ptr
	j := 0
	for i := 0; i < len(slice); i++ {
		if slice[i] != "" {
			slice[j] = slice[i]
			j++
		}
	}
	*ptr = slice[:j]
	return j
}
