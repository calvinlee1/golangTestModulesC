package c

import (
	"fmt"

	d "github.com/calvinlee1/golangTestModulesD"
)

func CallC() {
	fmt.Println("call C: v3.1.1")
	fmt.Println("   --> call D:")
	fmt.Printf("\t")
	d.CallD()
	fmt.Println("   --> call D end")
}
