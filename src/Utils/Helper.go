package Utils

import (
	"bytes"
	"strings"
)

func Join(str ...string) string {
	ret := ""
	for _, s := range str {
		ret += s
	}
	return ret
}

func GoJoin(strs ...string) string {
	return strings.Join(strs, "")
}

func BufferJoin(strs ...string) string {
	var bf bytes.Buffer
	for _, str := range strs {
		bf.WriteString(str)
	}
	return bf.String()
}
