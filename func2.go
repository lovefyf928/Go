package main

func test(x int) func() {
	return func() {
		println(x)
	}
}

func test2() {
	x := 100
	x ++
	println(x)
}

func test3() func() {
	x := 100
	return func() {
		x ++
		println(x)
	}
}

func main() {
	x := 100

	f := test(x)
	f()

	test2()
	test2()
	test2()

	f1 := test3()
	f1()
	f1()
	f1()
	f1()
	f1()
	f1()
}
