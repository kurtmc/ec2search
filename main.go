package main

import (
	"fmt"
	"os"

	"github.com/kurtmc/ec2search/search"
)

func main() {
	instances, err := search.ListInstances(os.Args[1])
	if err != nil {
		panic(err)
	}
	for _, i := range instances {
		fmt.Println(i)
	}
}
