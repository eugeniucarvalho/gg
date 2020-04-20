package generator

import (
	"encoding/json"
	"fmt"
	"strings"
	
	"git.eugeniocarvalho.dev/eugeniucarvalho/utils"
	G "github.com/dave/jennifer/jen"
)

const (
	// UTILS          = "git.gojus.com.br/eugeniucarvalho/utils"
	// GENERATOR_BASE = "git.gojus.com.br/eugeniucarvalho/gg/generator"
)

type Generator struct {
	Language string              `json:"language"`
	Keywords []string            `json:"keywords"`
	Commands map[string]*Command `json:"commands"`
	BuiltIn  map[string]*Command `json:"builtIn"`
	Output   string
}
type Command struct {
	Name               string  `json:"name"`
	IgnoreGen          bool    `json:"ignoreGen"`
	Delimiter          string  `json:"delimiter"`
	MethodName         string  `json:"-"`
	Template           string  `json:"template"`
	ParamList          []Param `json:"paramList,omitempty"`
	Description        string  `json:"description,omitempty"`
	RenderChildrenAs   string  `json:"readChildrenAs,omitempty"`
	ChildrenRenderMode string  `json:"childrenRenderMode,omitempty"`
}
type Param struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func New(file string) (*Generator, error) {
	data, err := utils.FileGetContents(file)
	if err == nil {
		g := &Generator{}
		if err = json.Unmarshal([]byte(data), g); err == nil {

			// g.BuiltInRegister()

			return g, nil
		}
	}
	return nil, err
}



func (g *Generator) AddCommand(key string, cmd *Command) {
	g.Keywords = append(g.Keywords, key)
	g.Commands[key] = cmd
}

func (g *Generator) Gen(folder string) (err error) {
	g.Output = folder
	file := G.NewFile(g.Language)

	// Cria interface com todos os metodos
	if err = g.GenBaseInterface(file); err != nil {
		return err
	}

	// Cria todas as funcoes staticas com os comandos basicos
	// g.GenBuiltIn(file)

	fmt.Printf("Filepath: %s\n", g.FilePath("struct.go"))

	if err = file.Save(g.FilePath(g.Language + ".go")); err != nil {
		panic(err)
	}

	return nil
}

func (g *Generator) FilePath(name string) string {
	return fmt.Sprintf("%s/%s/%s", g.Output, g.Language, name)
}

func (g *Generator) GenBaseInterface(file *G.File) (err error) {

	// file.Type().Id("File").StructFunc(func(g *G.Group) {
	// 	g.Add(G.Id("Group"))
	// 	g.Add(G.Id("Name").Id("string"))
	// }).Line()

	// file.Func().Id("NewFile").Params(G.Id("file").Id("string")).Params(G.Op("*").Id("File")).Block(
	// 	G.Return(G.Op("&").Id("File").Values(G.Dict{
	// 		G.Id("Name"): G.Id("file"),
	// 	})),
	// )

	// file.Func().Params(G.Id("f").Op("*").Id("File")).Id("Save").Params().Params(G.Id("error")).Block(
	// 	G.Return(G.Qual(UTILS, "FilePutContents").Call(
	// 		G.Id("f").Dot("Name"),
	// 		G.Id("f").Dot("GoString").Call(),
	// 		G.Lit(777),
	// 	)),
	// )

	// file.Comment("Entidade representa menor unidade de Statement.")

	// file.Type().Id("Stmt").StructFunc(func(g *G.Group) {
	// 	g.Add(G.Id("Group"))
	// 	g.Add(G.Id("value").Id("interface{}"))
	// 	g.Add(G.Id("template").Id("string"))
	// 	g.Add(G.Id("Separete").Id("string"))
	// 	// g.Add(G.Id("Childrens").Index().Id("CodeInterface"))
	// }).Line()

	// file.Func().Params(
	// 	G.Id("s").Op("*").Id("Stmt"),
	// ).Id("Render").Params(
	// 	G.Id("buffer").Op("*").Qual("bytes", "Buffer"),
	// 	// G.Id("buffer").Op("*").Id("bytes").Dot("Buffer"),
	// ).Params(
	// 	G.Id("err").Id("error"),
	// ).Block(
	// 	G.For(
	// 		G.List(G.Id("_"), G.Id("s")).Op(":=").Range().Id("s").Dot("Stmts"),
	// 	).Block(
	// 		G.Id("err").Op("=").Id("s").Dot("Render").Call(G.Id("buffer")),
	// 	).Line().Return(G.Empty()),
	// )

	// file.Comment("Root node group of Statement.")

	// file.Type().Id("Group").StructFunc(func(g *G.Group) {
	// 	g.Add(G.Qual(GENERATOR_BASE, "Group"))
	// 	// g.Add(G.Id("Stmts").Index().Id("CodeInterface"))
	// }).Line()

	// file.Comment("Metodo realiza a renderização de um grupo e todos os seus statements.")
	// file.Func().Params(
	// 	G.Id("g").Op("*").Id("Group"),
	// ).Id("Render").Params(
	// 	G.Id("buffer").Op("*").Qual("bytes", "Buffer"),
	// 	// G.Id("buffer").Op("*").Id("bytes").Dot("Buffer"),
	// ).Params(
	// 	G.Id("err").Id("error"),
	// ).Block(
	// 	G.For(
	// 		G.List(G.Id("_"), G.Id("s")).Op(":=").Range().Id("g").Dot("Stmts"),
	// 	).Block(
	// 		G.Id("err").Op("=").Id("s").Dot("Render").Call(G.Id("buffer")),
	// 	).Line().Return(G.Empty()),
	// )
	// file.Comment("Metodo Gera a string do arquivo.")
	// file.Func().Params(
	// 	G.Id("g").Op("*").Id("Group"),
	// ).Id("GoString").Params().Params(
	// 	G.Id("string"),
	// ).Block(
	// 	G.Id("buf").Op(":=").Qual("bytes", "Buffer").Values(),
	// 	// G.Id("buf").Op(":=").Id("bytes").Dot("Buffer"),
	// 	G.If(
	// 		G.Id("err").Op(":=").Id("g").Dot("Render").Call(G.Op("&").Id("buf")),
	// 		G.Id("err").Op("!=").Nil(),
	// 	).Block(
	// 		G.Id("panic").Call(G.Id("err")),
	// 	).Line().Return(G.Id("buf").Dot("String").Call()),
	// ).Line()

	// file.Func().Id("NewGroup").Params().Op("*").Id("Group").Block(
	// 	G.Return(G.Op("&").Id("Group").Values()),
	// ).Line()

	var (
		cmd    *Command
		uppKey string
		found  bool
		// interfacesMethods = []*G.Statement{}
		method  *G.Statement
		methods = []string{}
	)

	codeInterface := file.Type().Id("CodeInterface")
	interfaces := G.Statement{}
	// Adiciona a importacao da interface do gerador
	// interfaces = append(interfaces, G.Qual(GENERATOR_BASE, "CodeInterface"))
	interfaces = append(interfaces, G.Id("BaseCodeInterface"))

	// interfaces = append(interfaces, G.Id("Render").Params(
	// 	G.Id("buffer").Op("*").Qual("bytes", "Buffer"),
	// ).Params(G.Id("error")))

	for _, key := range g.Keywords {

		uppKey = strings.Title(key)

		methods = append(methods, uppKey)

		if cmd, found = g.Commands[key]; found {

			cmd.Name = key
			cmd.MethodName = uppKey

			interfaces = genCmd(file, cmd, interfaces)

		} else {
			// Todos os metodos para keywords genericas sem template
			// são tratados aqui.
			interfaces = append(interfaces, G.Id(uppKey).Params().Params(G.Op("*").Id("Group")))

			method = G.Func().Params(
				G.Id("g").Op("*").Id("Group"),
			).Id(uppKey).Params().Params(G.Op("*").Id("Group"))

			file.Add(method.Clone().Block(
				G.Id("s").Op(":=").Op("&").Id("Stmt").Values(G.Dict{
					G.Id("Value"):    G.Lit(key),
					G.Id("Template"): G.Lit(" %s "),
				}),

				G.Id("g").Dot("Stmts").Op("=").Append(
					G.Id("g").Dot("Stmts"),
					G.Id("s"),
				),
				G.Return(G.Id("g")),
			).Line())

			ModuleFunction(file, uppKey, G.Statement{}, []string{})
		}

	}

	codeInterface.Interface(interfaces...)

	return
}

func genCmd(file *G.File, cmd *Command, interfaces G.Statement) G.Statement {

	var (
		p           *G.Statement
		params      = G.Statement{}
		paramsKeys  = []string{}
		paramsTypes = []string{}
	)

	file.Func().Params(
		G.Id("g").Op("*").Id("Group"),
	).Id(cmd.MethodName).ParamsFunc(func(g *G.Group) {

		if len(cmd.ParamList) == 0 {
			return
		}

		for _, param := range cmd.ParamList {
			p = G.Id(param.Name).Id(param.Type)
			g.Add(p)
			paramsKeys = append(paramsKeys, param.Name)
			paramsTypes = append(paramsTypes, param.Type)
			params = append(params, p)
		}

		interfaces = append(interfaces, G.Id(cmd.MethodName).Params(params...).Params(G.Op("*").Id("Group")))

	}).Params(
		G.Op("*").Id("Group"),
	).BlockFunc(func(g *G.Group) {

		// g.Add(G.Id("s").Op(":=").Op("&").Qual(GENERATOR_BASE, "Stmt").ValuesFunc(func(x *G.Group) {
		g.Add(G.Id("s").Op(":=").Op("&").Id("Stmt").ValuesFunc(func(x *G.Group) {
			var (
				value interface{}
				// lit   bool
				lit = true
			)

			x.Add(G.Id("Template").Op(":").Lit(fmt.Sprintf(" %s ", cmd.Template)))

			for k, typ := range paramsTypes {
				// fmt.Println("---", k, typ, paramsKeys[k])

				switch {
				case typ == "string" || typ == "interface{}":

					value = paramsKeys[k]
					lit = false

				case typ[0:3] == "...":
					x.Add(
						// G.Id("Group").Op(":").Id("Group").Values(
						// G.Id("Group").Op(":").Qual(GENERATOR_BASE, "Group").Values(G.Dict{
						G.Id("Group").Op(":").Id("Group").Values(G.Dict{
							G.Id("Stmts"): G.Id(paramsKeys[k]),
						}),
					)
					paramsKeys[k] += "..."
				default:
					// x.Add(G.Id("Group").Op(":").Id("Group").Values(
					// 	G.Qual(GENERATOR_BASE, "Group").Op(":").Qual(GENERATOR_BASE, "Group").Values(G.Dict{
					// 		G.Id("Stmts"): G.Index().Id("CodeInterface").Values(G.Id(paramsKeys[k])),
					// 	}),
					// ))
					// x.Add(G.Id("Group").Op(":").Qual(GENERATOR_BASE, "Group").Values(G.Dict{
					// 	G.Id("Stmts"): G.Index().Qual(GENERATOR_BASE, "CodeInterface").Values(G.Id(paramsKeys[k])),
					// }))
					x.Add(G.Id("Group").Op(":").Id("Group").Values(G.Dict{
						G.Id("Stmts"): G.Index().Id("CodeInterface").Values(G.Id(paramsKeys[k])),
					}))
				}
				break
			}

			if value == nil {
				value = cmd.Name
			}

			if lit {
				x.Add(G.Id("Value").Op(":").Lit(value))
			} else {
				x.Add(G.Id("Value").Op(":").Id(value.(string)))
			}

			if cmd.Delimiter != "" {
				x.Add(G.Id("Delimiter").Op(":").Lit(cmd.Delimiter))
			}

		}))

		g.Add(G.Id("g").Dot("Stmts").Op("=").Append(
			G.Id("g").Dot("Stmts"),
			G.Id("s"),
		))
		g.Add(G.Return(G.Id("g")))

	}).Line()

	ModuleFunction(file, cmd.MethodName, params, paramsKeys)
	return interfaces

}

func ModuleFunction(file *G.File, method string, params G.Statement, paramsKeys []string) {

	file.Func().Id(method).Params(params...).Params(G.Op("*").Id("Group")).Block(

		G.Return(G.Id("NewGroup").Call().Dot(method).CallFunc(func(g *G.Group) {
			for _, key := range paramsKeys {
				g.Add(G.Id(key))
			}
		})),
	).Line()
}

// func (g *Generator) BuiltInRegister() {

// 	if g.BuiltIn == nil {
// 		g.BuiltIn = map[string]*Command{}
// 	}
// 	fmt.Println("Add built in")

// 	g.AddBuiltIn("params", &Command{
// 		Template: "(%s)",
// 		ParamList: []Param{
// 			Param{
// 				Name: "params",
// 				Type: "...CodeInterface",
// 			},
// 		},
// 		Description: "Parameter definition in function declaration.",
// 	})

// 	g.AddBuiltIn("op", &Command{
// 		Template: " %s ",
// 		ParamList: []Param{
// 			Param{
// 				Name: "op",
// 				Type: "string",
// 			},
// 		},
// 		Description: "Representa um operando {=,!=,>,<,>=,<=,-,+,*,/,%}.",
// 	})

// 	g.AddBuiltIn("block", &Command{
// 		// Template:         " {\n%s\n}\n",
// 		Template:         " {%s}\n",
// 		RenderChildrenAs: "lines",
// 		ParamList: []Param{
// 			Param{
// 				Name: "stmts",
// 				Type: "...CodeInterface",
// 			},
// 		},
// 	})

// 	g.AddBuiltIn("call", &Command{
// 		Template: "(%s)",
// 		ParamList: []Param{
// 			Param{
// 				Name: "params",
// 				Type: "CodeInterface",
// 			},
// 		},
// 		Description: "Generete a call of function or method. The parameter is um List Stmt.",
// 	})

// 	g.AddBuiltIn("id", &Command{
// 		Template: "%s",
// 		ParamList: []Param{
// 			Param{
// 				Name: "stmt",
// 				Type: "string",
// 			},
// 		},
// 	})

// 	g.AddBuiltIn("comment", &Command{
// 		Template:  "\n//%s\n",
// 		IgnoreGen: true,
// 		ParamList: []Param{
// 			Param{
// 				Name: "stmt",
// 				Type: "string",
// 			},
// 		},
// 	})

// 	g.AddBuiltIn("lit", &Command{
// 		Template:  "%s",
// 		IgnoreGen: true,
// 		ParamList: []Param{
// 			Param{
// 				Name: "stmt",
// 				Type: "interface{}",
// 			},
// 		},
// 	})

// 	g.AddBuiltIn("index", &Command{
// 		Template: "[%s]",
// 		ParamList: []Param{
// 			Param{
// 				Name: "index",
// 				Type: "CodeInterface",
// 			},
// 		},
// 		Description: "Gen a index access. The parameter is a List",
// 	})
// }

// func (g *Generator) AddBuiltIn(key string, cmd *Command) {
// 	// g.BuiltIn = append(g.BuiltIn, key)
// 	g.BuiltIn[key] = cmd
// }
// func (g *Generator) GenBuiltIn(file *G.File) {
// 	var (
// 		params     *G.Statement
// 		paramsKeys []string
// 		name       string
// 		typ        string
// 		ret        string
// 	)
// 	for method, cmd := range g.BuiltIn {

// 		paramsKeys = []string{}
// 		method = strings.Title(method)

// 		params = &G.Statement{}

// 		for _, param := range cmd.ParamList {

// 			name = param.Name
// 			typ = param.Type
// 			ret = ""
// 			if strings.Contains(typ, "...") {
// 				ret = "..."
// 				name += ret
// 				typ = typ[3:]
// 			}

// 			if strings.Contains(param.Type, "CodeInterface") {

// 				params.Add(G.Id(param.Name).Op(ret).Qual(GENERATOR_BASE, typ))
// 			} else {
// 				params.Add(G.Id(param.Name).Id(typ))

// 			}

// 			paramsKeys = append(paramsKeys, name)

// 			// params = append(params, G.Id(param.Name).Id(param.Type))
// 		}

// 		file.Func().Params(
// 			G.Id("g").Op("*").Id("Group"),
// 		).Id(method).Params(*params...).Params(
// 			G.Op("*").Id("Group"),
// 		).Block(

// 			// G.Id("g").Dot(method).CallFunc(func(g *G.Group) {
// 			// 	for _, key := range paramsKeys {
// 			// 		g.Add(G.Id(key))
// 			// 	}
// 			// }),
// 			G.Id("g").Dot("Group").Dot(method).CallFunc(func(g *G.Group) {
// 				for _, key := range paramsKeys {
// 					g.Add(G.Id(key))
// 				}
// 			}),

// 			G.Return(G.Id("g")),
// 		).Line()

// 		file.Func().Id(method).Params(*params...).Params(
// 			G.Op("*").Id("Group"),
// 		).Block(
// 			G.Return(G.Id("NewGroup").Call().Dot(method).CallFunc(func(g *G.Group) {
// 				for _, key := range paramsKeys {
// 					g.Add(G.Id(key))
// 				}
// 			})),
// 		).Line()
// 	}
// }
