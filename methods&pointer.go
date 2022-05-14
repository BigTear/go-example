package main

type rect struct {
	width, height int
}

func (this *rect) area() int {
	return this.width * this.height
}

func (this rect) perim() int {
	return 2*this.width + 2*this.height
}

func main() {
	r := rect{width: 10, height: 5}
	println("area:", r.area())
	println("perim:", r.perim())

	rp := &r
	println("area:", rp.area())
	println("perim:", rp.perim())

}
