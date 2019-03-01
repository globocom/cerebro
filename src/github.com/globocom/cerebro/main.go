package main

import (
	"github.com/globocom/cerebro/modules"
)

func main() {
	modules.Init(modules.NewESClient)
}
