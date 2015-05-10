package main

import (
	"fmt"
	"os"
)

func fatal(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func checkError(err error) {
	if err != nil {
		//panic(err)
		fmt.Printf("[ERR] %s\n", err)
		os.Exit(1)
	}
}
