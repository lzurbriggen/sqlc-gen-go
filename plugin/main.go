package main

import (
	"github.com/sqlc-dev/plugin-sdk-go/codegen"

	golang "github.com/sqlc-dev/sqlc-gen-go/internal"
)

var Generate = golang.Generate

func main() {
	codegen.Run(golang.Generate)
}
