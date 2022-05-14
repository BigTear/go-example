package main

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	netInt := intSeq()

	println(netInt())
	println(netInt())
	println(netInt())
	println()

	newInt := intSeq()
	println(newInt())
	println(newInt())
	println(newInt())
	println()

	println(intSeq())
	println(intSeq())
	println(intSeq())
}
