package main

func main() {
	x := 4

	for {
		println(x)
		x--

		if x < 0 {
			break
		}
	}
}
