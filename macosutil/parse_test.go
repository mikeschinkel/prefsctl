package macosutil_test

import (
	"testing"

	"github.com/mikeschinkel/prefsctl/macosutil"
)

func TestFloatParser(t *testing.T) {
	validCases := []struct {
		prec int
		val  string
	}{
		{0, "0"},
		{0, "42"},
		{0, "-42"},
		{1, "0.0"},
		{0, "1."},
		{5, "3.14159"},
		{5, "-3.14159"},
		{1, ".5"},
		{1, "-.5"},
		{2, "1.23"},
		{2, "+1.23"},
		{3, "123.456"},
		{3, "-123.456"},
		{7, "0.0000001"},
		{7, "-0.0000001"},
	}

	invalidCases := []string{
		"",
		".",
		"-",
		"abc",
		"1.2.3",
		"-.",
		"e5",
		"1e5",     // Scientific notation not included
		"-1e-10",  // Scientific notation not included
		"0x1.2p3", // Hex float not included
	}

	for _, tc := range validCases {
		prec, err := macosutil.GetFloatPrecision(tc.val)
		if err != nil {
			t.Errorf("Failed to parse %q; %v", tc.val, err)
		}
		if prec != tc.prec {
			t.Errorf("Expected %q to match pattern", tc.val)
		}
	}

	for _, tc := range invalidCases {
		_, err := macosutil.GetFloatPrecision(tc)
		if err == nil {
			t.Errorf("Expected %q to NOT match pattern", tc)
		}
	}
}

func TestGetFloatPrecision(t *testing.T) {
	tests := []struct {
		value    string
		wantPrec int
		wantErr  bool
	}{
		{
			value:    "1.23foobar",
			wantPrec: 2,
			wantErr:  true,
		},
		{
			value:    "1.23",
			wantPrec: 2,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			gotPrec, err := macosutil.GetFloatPrecision(tt.value)
			if tt.wantErr {
				if err != nil {
					return
				}
				t.Errorf("GetFloatPrecision() error = %v, wantErr %v", err, tt.wantErr)
			}
			if gotPrec != tt.wantPrec {
				t.Errorf("GetFloatPrecision() gotPrec = %v, want %v", gotPrec, tt.wantPrec)
			}
		})
	}
}
