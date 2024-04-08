package main

import "log"

func main() {
	r := wireGin()
	log.Fatalln(r.Run())
}
