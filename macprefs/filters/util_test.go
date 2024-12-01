package filters

import (
	"testing"
)

func TestCodify(t *testing.T) {
	tests := []struct {
		name string
		args string
		want Code
	}{
		{"All lowercase w/o space", "abcdef", "abcdef"},
		{"All lowercase w/a space", "abc def", "abcDef"},
		{"All uppercase w/a space", "ABCDEF", "aBCDEF"},
		{"Proper case", "AbcDef", "abcDef"},
		{"Camelcase", "abcDef", "abcDef"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Codify(tt.args); got != tt.want {
				t.Errorf("Codify() = %v, want %v", got, tt.want)
			}
		})
	}
}
