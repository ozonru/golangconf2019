package calc

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
	"strconv"
)

func CalcRWords(txt, letter string) int {
	if runtime.GOOS == "windows" {
		panic("WRONG OS, TRY USING UNIX")
	}

	var buf bytes.Buffer
	buf.Write([]byte(txt))

	cmd := fmt.Sprintf(`for (/(^|\s)%s/g) {$a++;} END {print $a+0}`, letter)
	perl := exec.Command("perl", "-ne", cmd)
	perl.Stdin = &buf
	res, err := perl.CombinedOutput()
	if err != nil {
		panic(fmt.Sprintf("Error in calculation: %v", err))
	}

	intRes, err := strconv.Atoi(string(res))
	if err != nil {
		panic(fmt.Sprintf("Error in convertion: %v", err))
	}

	return intRes
}
