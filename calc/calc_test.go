package calc

import (
	"testing"
)

func TestGetTargetLetterStartWords(t *testing.T) {
	type args struct {
		txt    string
		letter string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "test 1",
			args: args{
				txt:    "ааа ррр ббб",
				letter: "р",
			},
			want: 1,
		},
		{
			name: "test 2",
			args: args{
				txt:    "ааа ррр ббб\nробот\nработа твоя",
				letter: "р",
			},
			want: 3,
		},
		{
			name: "test 3",
			args: args{
				txt: `
Главное, любой бы засмущался, находясь
на вечеринке в штанах на пять размеров больше. А
этому – все по-барабану. Взял еще пива, попросил
бумажку с ручкой и стал что-то быстро писать на
листочке.
`,
				letter: "р",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTargetLetterStartWords(tt.args.txt, tt.args.letter); got != tt.want {
				t.Errorf("getTargetLetterStartWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
