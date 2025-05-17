package main

import "fmt"

func main() {
	str := "01000101100"
	sikz := len(str)/8 + min(len(str)%8, 1)
	byt := make([]byte, sikz)

	now := 0
	ok := uint8(0)
	idx := 0
	for i := len(str) - 1; i >= 0; i -= 1 {

		if str[i] == '1' {
			ok += 1 << now
		}

		fmt.Printf("%v\n", ok)

		if now == 7 {
			byt[idx] = ok
			idx++
			now = 0
			ok = 0
		}
		now++
	}

	if now != 1 && ok > 0 {
		fmt.Printf("%v\n", now)
		byt[idx] = ok
	}

	fmt.Printf("%v and %v\n", byt, string(byt))
	for _, v := range string(byt) {
		fmt.Printf("%v\n", rune(v))
	}
	fmt.Printf("%v", len(byt))
}
