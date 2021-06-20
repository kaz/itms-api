package main

import (
	"github.com/kaz/itms-api/internal/server"
)

func main() {
	if err := server.Start(); err != nil {
		panic(err)
	}
}
