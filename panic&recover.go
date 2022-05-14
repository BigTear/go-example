package main

func 可能恐慌() {
	panic("aaaaa problem")
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			println("recovered. error:\n", r)
		}
	}()

	可能恐慌()

	println("after panic")

}
