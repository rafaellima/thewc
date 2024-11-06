package wc

import (
	"strings"
	"unicode/utf8"
)

type Wc struct {
	Data string
}

func (wc *Wc) CountChars() int {
	return utf8.RuneCountInString(string(wc.Data))
}

func (wc *Wc) CountBytes() int {
	return len(wc.Data)
}

func (wc *Wc) CountLines() int {
	return strings.Count(string(wc.Data), "\n")
}

func (wc *Wc) CountWords() int {
	return len(strings.Fields(string(wc.Data)))
}
