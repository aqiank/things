package main

import (
	"fmt"

	"../.."
)

const token = "eUMTupT1lr-OlQnnTkS22hhBCn2ygxH9hI45qAFv2Pg"

func main() {
	var s string
	var err error

	t := things.Thing{token}

	if s, err = t.ReadValue("foo"); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(s)
}
