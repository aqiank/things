package main

import (
	"fmt"

	"../.."
)

const token = "eUMTupT1lr-OlQnnTkS22hhBCn2ygxH9hI45qAFv2Pg"

func handler(keyValues []things.KeyValue) {
	fmt.Println(keyValues)
}

func main() {
	t := things.Thing{token}

	if _, err := t.Subscribe(handler); err != nil {
		fmt.Println(err)
		return
	}
	
	select{}
}
