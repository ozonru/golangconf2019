package calc

import (
	"testing"
	"gotest.tools/assert"
	)

func TestCalcRWords(t *testing.T) {
	assert.Equal(t, CalcRWords("ракета"), 1)
	assert.Equal(t, CalcRWords("ааа ррр ббб"), 1)
	assert.Equal(t, CalcRWords("ааа ррр ббб\nробот\nтвоя работа"), 3)
	assert.Equal(t, CalcRWords(`
Главное, любой бы засмущался, находясь
на вечеринке в штанах на пять размеров больше. А
этому – все по-барабану. Взял еще пива, попросил
бумажку с ручкой и стал что-то быстро писать на
листочке.
`), 2)
}