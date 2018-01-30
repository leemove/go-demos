package main

import (
	"fmt"
)

func main() {
	str := "12314512312398798710982309182309182300099890"
	str = fmtStr(str)
	fmt.Print(str)
}

func fmtStr(s string) string {

	if l := len(s); l > 3 {
		return fmtStr(s[:l-3]) + "," + s[l-3:]
	}
	return s
}
