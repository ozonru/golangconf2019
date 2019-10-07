package calc

import (
	"testing"

	"gotest.tools/assert"
)

func TestCalcRWords(t *testing.T) {
	assert.Equal(t, CalcRWords("ааа ррр ббб ррр", "р"), 2)
	assert.Equal(t, CalcRWords("ааа ррр ббб\nробот\nработа твоя\nработа", "р"), 4)
	assert.Equal(t, CalcRWords(`
Главное, любой бы засмущался, находясь
на вечеринке в штанах на пять размеров больше. А
этому – все по-барабану. Взял еще пива, попросил
бумажку с ручкой и стал что-то быстро писать на
листочке.
`, "р"), 2)
}
