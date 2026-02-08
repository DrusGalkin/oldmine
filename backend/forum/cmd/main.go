package main

import (
	"fmt"
	"forum/internal/config"
	"github.com/DrusGalkin/libs"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
	libs.LoggerInit(cfg.Env)
}
