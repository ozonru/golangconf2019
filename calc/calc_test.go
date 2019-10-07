package calc

import (
	"testing"

	"gotest.tools/assert"
)

func TestCalcRWords(t *testing.T) {
	tt := []struct {
		name        string
		in          string
		search      string
		shouldPanic bool
		expect      int
	}{
		{
			name:   "lower only",
			in:     "ааа ррр ббб",
			search: "р",
			expect: 1,
		},
		{
			name:   "работа",
			in:     "работа",
			search: "р",
			expect: 1,
		},
		{
			name:   "Работа",
			in:     "Работа",
			search: "р",
			expect: 1,
		},
		{
			name: "capitals in text",
			in: `
Ехал Грека через реку.
Видит Грека в реке Рак.
Сунул в реку руку Грека.
Рак за руку Греку - цап.
`,
			search: "р",
			expect: 7,
		},
		{
			name: "search capital",
			in: `
Ехал Грека через реку.
Видит Грека в реке Рак.
Сунул в реку руку Грека.
Рак за руку Греку - цап.
`,
			search: "Р",
			expect: 7,
		},
		{
			name:   "trim spaces",
			in:     "   ааа ррр ббб  ляля  ",
			search: "р",
			expect: 1,
		},
		{
			name:   "single letter words",
			in:     "аа а р рр б бб  ля ля р",
			search: "р",
			expect: 3,
		},
		{
			name:   "new lines",
			in:     "ааа ррр ббб\nробот\nработа твоя",
			search: "р",
			expect: 3,
		},
		{
			name:        "not single search",
			in:          "ааа ррр ббб\nробот\nработа твоя",
			search:      "ха",
			shouldPanic: true,
		},
		{
			name:        "not letter",
			in:          "ааа ррр ббб\nробот\nработа твоя",
			search:      "6",
			shouldPanic: true,
		},
		{
			name: "multiline",
			in: `
Главное, любой бы засмущался, находясь
на вечеринке в штанах на пять размеров больше. А
этому – все по-барабану. Взял еще пива, попросил
бумажку с ручкой и стал что-то быстро писать на
листочке.
`,
			search: "р",
			expect: 2,
		},
		{
			name: "multilang",
			in: `
Главное, любой this is some english words
бы засмущался, находясь
на вечеринке в штанах на пять размеров больше. А
этому – все по-барабану.
Chineese? 世界! Pretty Woman is my favourite song...
Взял еще пива, попросил
бумажку с ручкой и стал что-то быстро писать на
листочке.
`,
			search: "р",
			expect: 2,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if tc.shouldPanic {
				defer func() {
					err := recover()
					assert.Equal(t, err.(string), "invalid search letter")
				}()
			}
			actual := CalcRWords(tc.in, tc.search)
			assert.Equal(t, tc.expect, actual)
		})
	}
}

// old: BenchmarkCalcRWords-4   	  100000	     17916 ns/op	    1952 B/op	     278 allocs/op
// new: BenchmarkCalcRWords-4   	  500000	      3120 ns/op	       0 B/op	       0 allocs/op
func BenchmarkCalcRWords(b *testing.B) {
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
