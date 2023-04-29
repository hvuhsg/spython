def fib(n: int) -> int:
	if n < 1:
		return 1
	
	return fib(n-1) + fib(n+1)

a = fib(10)
