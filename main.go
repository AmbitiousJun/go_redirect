package main

import (
	"go_redirect/internal/config"
	"go_redirect/internal/pathresv"
	"go_redirect/internal/web"
)

func main() {
	if err := config.Init(); err != nil {
		panic(err)
	}
	pathresv.InitTemplates()
	if err := web.Listen(); err != nil {
		panic(err)
	}
}
