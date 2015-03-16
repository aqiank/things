package main

import (
	"fmt"

	"../.."
)

const code = "K0M_VXsT4mKC63IthuVGeLV_7Cfq1NJWhQ"

func main() {
	if _, err := things.Activate(code); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successfully activated the thing!")
}
