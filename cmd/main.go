package main

import (
	"log"

	"github.com/codern-org/gateway/pkg/infrastructure"
)

func main() {
	if err := infrastructure.OpenFiberServer(); err != nil {
		log.Fatal(err)
	}
}
