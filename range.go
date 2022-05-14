package main

func main() {
	kvs := map[string]string{"one": "1st", "two": "2nd"}
	for k, v := range kvs {
		println(k, v)
	}
	for k := range kvs {
		println(k)
	}
	for i := range "golang" {
		println(i)
	}
}
