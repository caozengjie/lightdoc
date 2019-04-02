package main

import (
	"lightdoc/config"
	"lightdoc/route"
	"os"
)

func main() {

	if len(os.Args) > 1 {
		config.Root = os.Args[1]
	}
	err := route.Init()
	if err != nil {
		panic(err)
	}
}