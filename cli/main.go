package main

import (
	"os"

	handler "github.com/kaz/itms-api"
)

func main() {
	if err := handler.HandleControl(os.Args[1], os.Args[2]); err != nil {
		panic(err)
	}
}
