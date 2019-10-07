package calc

import (
	"testing"
	"gotest.tools/assert"
)

func TestCalcRWords(t *testing.T) {
	assert.Equal(t, CalcWordsStartedWithLetter("ааа ррр ббб", "р"), 1)
	assert.Equal(t, CalcWordsStartedWithLetter("ааа ррр ббб\nробот\nработа твоя", "р"), 3)
	assert.Equal(t, CalcWordsStartedWithLetter(`
Главное, любой бы засмущался, находясь
на вечеринке в штанах на пять размеров больше. А
этому – все по-барабану. Взял еще пива, попросил
бумажку с ручкой и стал что-то быстро писать на
листочке.
`, "р"), 2)
	assert.Equal(t, CalcWordsStartedWithLetter("Тест на ошибку, когда слово на 'р' в кавычках", "р"), 1)
	assert.Equal(t, CalcWordsStartedWithLetter("тест на Разный case слова на р", "р"), 2)
}