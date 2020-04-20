package main

import (
	"flag"
	"git.eugeniocarvalho.dev/eugeniucarvalho/gg/generator"
)

var (
	Input = flag.String("input", "def.json", "Arquivo de descrição do projeto no formato JSON.")
	Out   = flag.String("out", "./generators", "Destino dos arquivos gerados.")
)

//a.replace(/[\n\r]/g,";").replace(/\s+/g,'').split(";").map(function(k,v){dic[k]=1}); console.log(Object.getOwnPropertyNames(dic).sort())

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
