package main

import (
	"fmt"

	"../.."
)

const token = "eUMTupT1lr-OlQnnTkS22hhBCn2ygxH9hI45qAFv2Pg"

func main() {
	var err error

	t := things.Thing{token}
	
	if err = t.WriteValue("foo", "42"); err != nil {
		fmt.Println(err)
		return
	}
}
