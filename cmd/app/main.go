package main

import "wish/internal/app"

const configsDir = "configs"

func main() {
	app.Run(configsDir)
}
