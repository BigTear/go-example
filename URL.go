package main

import "net/url"

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	println("URL:", s)

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	println("Scheme:", u.Scheme)
	println(u.User, u.User.Username())
}
