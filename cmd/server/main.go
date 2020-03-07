package main

import (
	"github.com/naoto67/clean-architecture/src/config"
	"github.com/naoto67/clean-architecture/src/interface/api/router"
)

func main() {
	config.Setup()

	r := router.New()
	r.Run()
}
