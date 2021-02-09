package main

import (
	"flag"

	"github.com/eugeniucarvalho/gg/generator"
)

var (
	Input = flag.String("input", "def.json", "Arquivo de descrição do projeto no formato JSON.")
	Out   = flag.String("out", "./generators", "Destino dos arquivos gerados.")
)

func main() {
	var (
		err         error
		description *generator.Generator
	)

	flag.Parse()

	if description, err = generator.New(*Input); err != nil {
		panic(err)
	}

	description.Gen(*Out)
}
