package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	const s = "สวัสดี"
	fmt.Println(s, len(s))

	fmt.Print("以%x格式输出:")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	fmt.Println("\nRune 计数:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\n使用 DecodeRuneInString")

	var examineRune func(r rune)
	examineRune = func(r rune) {
		if r == 't' {
			fmt.Println("tee")
		} else if r == 'ด' {
			fmt.Println("so sua")
		}
	}

	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}
