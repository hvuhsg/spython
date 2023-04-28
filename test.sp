n = 10000000
a = 0
b = 1

while n > 0:
	n = n - 1
	b = a + b
	a = b - a
