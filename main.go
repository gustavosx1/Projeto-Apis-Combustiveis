package main

import (
	"go-cli/api"
)

func main() {
	URL := "https://combustivelapi.com.br/api/precos/"
	api.Request(URL, "a", "gasolina")
}
