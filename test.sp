def fib(n: int) -> int:
	a = 0
	b = 1

	while n > 0:
		n = n - 1
		b = a + b
		a = b - a
	
	return b

a = fib(41)
return a