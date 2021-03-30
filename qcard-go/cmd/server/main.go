package main

import (
	"os"

	"github.com/nucktwillieren/project-d/qcard-go/internal/api"
)

func main() {
	args := os.Args
	r := api.Setup()
	//r.RunTLS(
	//	":8080",
	//	"./fullchain.pem",
	//	"./privkey.pem",
	//)

	if len(args) > 1 {
		r.Run(":" + args[1])
	} else {
		r.Run()
	}
}
