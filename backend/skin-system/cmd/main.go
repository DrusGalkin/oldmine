package main

import (
	"libs"
	"skin_system/internal/app"
	"skin_system/internal/config"
	"skin_system/pkg"
)

const UPLOAD_PATH = "./uploads"

func main() {
	pkg.MustLoadMkDir(UPLOAD_PATH)

	cfg := config.MustLoad()
	log := libs.LoggerInit(cfg.Env)

	app.Run()

}
