package simple

func show(n int) {
	if n > 5 {
		show(n - 1) //开辟栈
	}
	println(n)
}

//阶乘
func factorial(n int) int {
	if n == 1 {
		return 1
	} else {
		return factorial(n-1) * n // n-1 * n
	}
}
