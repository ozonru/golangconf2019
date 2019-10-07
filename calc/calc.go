package calc

import (
	"regexp"
	"sync"
)

// \todo remove global state
var (
	once sync.Once
	re   *regexp.Regexp
)

// \todo rename this go-way & change func signature to remove global state
func CalcRWords(txt string, letter string) int {
	once.Do(func() {
		expr := `(?:\A|\z|\s)` + letter + `+`
		re = regexp.MustCompile(expr)
	})
	res := re.FindAllStringSubmatchIndex(txt, -1)
	return len(res)
}
