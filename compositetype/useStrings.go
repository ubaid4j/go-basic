package main

import (
	"fmt"
	s "strings"
	"unicode"
)

var f = fmt.Printf

func main() {
	upper := s.ToUpper("Hello there!")
	f("To Upper: %s\n", upper)
	f("To Lower: %s\n", s.ToLower("Hello THERE"))
	f("%s\n", s.Title("tHis will be a title!"))
	f("EqualFold: %v\n", s.EqualFold("ubaid", "UBAiD"))
	f("EqualFold: %v\n", s.EqualFold("ubaid", "UBAI"))

	f("Prefix: %v\n", s.HasPrefix("Ubaid", "Ub"))
	f("Prefix: %v\n", s.HasPrefix("Ubaid", "ub"))
	f("Suffix: %v\n", s.HasSuffix("Ubaid", "id"))
	f("Suffix: %v\n", s.HasSuffix("Ubaid", "ID"))

	f("Index: %v\n", s.Index("Ubaid", "ai"))
	f("Index: %v\n", s.Index("Ubaid", "Ai"))
	f("Count: %v\n", s.Count("Ubaid", "i"))
	f("Count: %v\n", s.Count("Ubaid", "I"))
	f("Repeat: %s\n", s.Repeat("ab", 5))

	f("TrimSpace: %s\n", s.TrimSpace("   \tThis is a line. \n"))
	f("TrimLeft: %s", s.TrimLeft("   \tThis is a\t line. \n", "\n\t "))
	f("TrimRight: %s\n", s.TrimRight("  \tThis ia a\t line. \n", "\n\t "))

	f("Compare: %v\n", s.Compare("Ubaid", "UBAID"))
	f("Compare %v\n", s.Compare("Ubaid", "Ubaid"))
	f("Compare: %v\n", s.Compare("UBAID", "UBAid"))

	f("Fields: %v\n", s.Fields("This is a string!"))
	f("Fields: %v\n", s.Fields("Thisis\na\tstring!"))

	f("%s\n", s.Split("abcd efg", ""))

	f("%s\n", s.Replace("abcd efg", "", "_", -1))
	f("%s\n", s.Replace("abcd efg", "", "_", 4))
	f("%s\n", s.Replace("abcd efg", "", "_", 2))

	lines := []string{"Line 1", "Line 2", "Line 3"}
	f("Join: %s\n", s.Join(lines, "+++"))
	f("SplitAfter: %s\n", s.SplitAfter("123+++432++", "++"))
	trimFunction := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	f("TrimFunc: %s\n", s.TrimFunc("123 abc ABC \t", trimFunction))

}
