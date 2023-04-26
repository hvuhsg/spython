define i64 @main() {
0:
	%a = alloca i64
	store i64 5, i64* %a
	%1 = ptrtoint i64* %a to i64
	%2 = icmp eq i64 %1, 5
	br i1 %2, label %4, label %5

3:
	ret i64 0

4:
	store i64 0, i64* %a
	br label %3

5:
	store i64 1, i64* %a
	br label %3
}

