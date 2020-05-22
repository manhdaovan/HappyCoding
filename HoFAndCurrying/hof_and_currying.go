package hofandcurrying

import (
	"fmt"
	"strings"
)

// SelectStrings returns strings with length over 3
func SelectStrings(strs []string) []string {
	if len(strs) == 0 {
		return nil
	}
	ret := make([]string, 0, len(strs))
	for _, s := range strs{
		if len(s) > 3 {
			ret = append(ret, s)
		}
	}

	return ret
}

type ChooseFnc func(s string) bool

func LengthOver(n uint) ChooseFnc {
	return func(s string) bool {
		return uint(len(s)) > n
	}
}

func StartsWith(c string) ChooseFnc {
	return func(s string) bool {
		return strings.HasPrefix(s, c)
	}
}

func SelectStringsHof(strs []string, f ChooseFnc) []string {
	if len(strs) == 0 {
		return nil
	}

	ret := make([]string, 0, len(strs))
	for _, s := range strs{
		if f(s) {
			ret = append(ret, s)
		}
	}

	return ret
}

func UsingSelectStrings() {
	strs := []string{"a", "ab", "bbc", "dbcd", "abcde"}
	strs1 := SelectStringsHof(strs, LengthOver(5))
	fmt.Println("strs1 -- ", strs1)
	strs2 := SelectStringsHof(strs, StartsWith("b"))
	fmt.Println("strs2 -- ", strs2)
}

