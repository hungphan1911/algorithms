package test

import (
	"github.com/hungphan1911/algorithms/src/math/numbertheory"
	"testing"
)

func TestBinaryExponentiation(t *testing.T) {
	tests := []struct {
		n, p, m uint
		want    uint
	}{
		{2, 3, 5, 3},  // 2^3 = 8, 8 % 5 = 3
		{3, 4, 7, 4},  // 3^4 = 81, 81 % 7 = 4
		{10, 0, 7, 1}, // anything^0 = 1
		{5, 5, 13, 5}, // 5^5 = 3125, 3125 % 13 = 5
	}

	for _, tt := range tests {
		got := numbertheory.BinaryExponentiation(tt.n, tt.p, tt.m)
		if got != tt.want {
			t.Errorf("BinaryExponentiation(%d, %d, %d) = %d; want %d", tt.n, tt.p, tt.m, got, tt.want)
		}
	}
}
