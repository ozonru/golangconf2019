package calc

import (
	"testing"

	"gotest.tools/assert"
)

func TestCalcRWords(t *testing.T) {
	assert.Equal(t, CalcRWords("ааа ррр ббб  р", "р"), 2)
	assert.Equal(t, CalcRWords("ааа ррр ббб\nробот\nработа твоя", "р"), 3)
	assert.Equal(t, CalcRWords(`
Главное, любой бы засмущался, находясь
на вечеринке в штанах на пять размеров больше. А
этому – все по-барабану. Взял еще пива, попросил
бумажку с ручкой и стал что-то быстро писать на
листочке.
`, "р"), 2)
	assert.Equal(t, CalcRWords(`С использованием
Капсового капса`, "к"), 2)
	assert.Equal(t, CalcRWords("'расческа', \"Расчесочка\", \\{расческа}\\", "р"), 3)
}

func BenchmarkCalcRWords(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		CalcRWords(`
Главное, любой бы засмущался, находясь
на вечеринке в штанах на пять размеров больше. А
этому – все по-барабану. Взял еще пива, попросил
бумажку с ручкой и стал что-то быстро писать на
листочке.
`, "р")
	}
}
