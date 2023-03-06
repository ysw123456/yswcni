package main

import (
	"cni/utils"
	"fmt"
)

func main() {

	fmt.Println("Hello GO!!!")
	utils.Write("./test", "first log %s", "hello world")
}
