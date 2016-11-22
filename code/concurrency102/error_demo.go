package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	if contents, err := ioutil.ReadFile("/etc/hosts"); err != nil {
		fmt.Println(err.Error())
		fmt.Printf("%#v\n", err)
	} else {
		fmt.Printf("%s", contents)
	}
}
