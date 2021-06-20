package main

import (
	"os"

	"github.com/kaz/itms-api/internal/server"
)

func main() {
	if err := server.Start(os.Getenv("K_SERVICE") == ""); err != nil {
		panic(err)
	}
}
