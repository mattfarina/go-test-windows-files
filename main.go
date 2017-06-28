package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.RemoveAll("files")
	fmt.Println(err)
}
