package main

import "fmt"

func main() {
	adoc := struct{ name string }{name: "tom"}
	fmt.Printf(adoc.name)
	bdoc := adoc
	bdoc.name = "jom"
	fmt.Println(bdoc)
}
