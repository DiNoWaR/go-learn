package main

import "fmt"

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func checkErr(err error) {
	fmt.Println(err == nil)
}

func main() {
	var err error
	checkErr(err)

	var e *errorString
	checkErr(e)

	e = &errorString{}
	checkErr(e)

	e = nil
	checkErr(e)
}
