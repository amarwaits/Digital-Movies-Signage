package main

import "github.com/amarwaits/Syndica-THT/src/app"

func main() {
	app := app.NewApp()

	err := app.Start()
	if err != nil {
		panic(err)
	}
}
