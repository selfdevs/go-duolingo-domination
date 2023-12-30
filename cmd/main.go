package main

import (
	duolingo "duolingo"
	"github.com/scaleway/serverless-functions-go/local"
)

func main() {
	local.ServeHandler(duolingo.Handle, local.WithPort(8080))
}
