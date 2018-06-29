package main

import (
	"fmt"
	"os"
	"unicode"

	"golang.org/x/exp/mmap"
)

var BasicLatin = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x0020, 0x007f, 1},
	},
	LatinOffset: 6,
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: gostrings {file}")
		return
	}

	p := os.Args[1]

	file, err := mmap.Open(p)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	offset := 0

	s := ""

	for i := 0; i < file.Len(); i++ {
		b := file.At(i)

		c := rune(b)

		if unicode.Is(BasicLatin, c) {
			s += fmt.Sprintf("%c", c)
			continue
		} else if c == 0x0d {
			s += fmt.Sprintf("%c", c)
			continue
		} else if c == 0x09 {
			s += fmt.Sprintf("%c", c)
			continue
		} else if c == 0x0a {
			s += fmt.Sprintf("%c", c)
			continue
		} else if len(s) > 6 {
			fmt.Printf("0x%x: %s\n", offset, s)

			s = ""
			offset = i + 1

			continue
		} else {
			s = ""
			offset = i + 1
			continue
		}

		/*
			// only ascii supported now
			if unicode.IsLetter(rune(b)) {
			} else if unicode.IsDigit(rune(b)) {
			} else if unicode.IsPunct(rune(b)) {
			} else if unicode.IsSpace(rune(b)) {
			} else if unicode.IsGraphic(rune(b)) {
				s = ""
				offset = i + 1
			} else if len(s) > 6 {
				fmt.Printf("0x%x: %s\n", offset, s)

				s = ""
				offset = i + 1
			} else {
				s = ""
				offset = i + 1
				continue
			}
		*/

	}
}
