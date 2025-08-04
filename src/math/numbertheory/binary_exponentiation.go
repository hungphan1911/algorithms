package numbertheory

// Iterative implementation of binary exponentiation
func BinaryExponentiation(n, p, m uint) uint {
	n %= m
	res := uint(1)

	for p > 0 {
		if p&1 != 0 {
			res = res * n % m
		}
		n = n * n % m
		p = p >> 1
	}

	return res
}
