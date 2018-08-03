package main

import (
	"fmt"
	"github.com/root-secure/showdon/libdon"
)

func main() {
	don := libdon.NewDon()
	fmt.Println(don.Get())
}
