package Utils

func Join(str ...string) string {
	ret := ""
	for _, s := range str {
		ret += s
	}
	return ret
}
