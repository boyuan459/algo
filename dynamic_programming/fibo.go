package dynamic_programming

func fibo(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return fibo(n-2) + fibo(n-1)
}

// the ideal of dynamic programming is to save the subresult to avoid recompute

func fibo_dp(n int) int {
	var f0 = 0
	var f1 = 1
	var f2 = 0

	if n == 1 {
		return 1
	}

	for i := 2; i <= n; i++ {
		f2 = f0 + f1
		f0 = f1
		f1 = f2
	}

	return f2
}
