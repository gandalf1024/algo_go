package simple

func show(n int) {
	if n > 5 {
		show(n - 1)
	}
	println(n)
}
