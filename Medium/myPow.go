package main

func myPow(x float64, n int) float64 {
	arr := []float64{x}

	if x == 1 || n == 0 || x == -1 && n%2 == 0 {
		return 1
	}
	if (x == -1 && n%2 != 0) || (x == -1 && n%2 == 0) {
		return -1
	}
	if x != 1 && n > 0 {
		for i := 1; i < n; i++ {
			temp := arr[0]
			x *= temp
		}
		return x
	}
	if x != 1 && n < 0 {
		for i := 1; i < (-1 * n); i++ {
			temp := arr[0]
			x *= temp
		}
		return 1 / x
	}
	return x
}
