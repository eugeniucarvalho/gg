package typescript

import (
	bytes "bytes"
	"fmt"
	"os/exec"
	"reflect"
	"strings"

	"git.eugeniocarvalho.dev/eugeniucarvalho/utils"
)

var (
	LitTemplates = map[string]string{
		"string":  " '%s' ",
		"int":     " %d ",
		"int8":    " %d ",
		"int16":   " %d ",
		"int32":   " %d ",
		"int64":   " %d ",
		"float32": " %.2f ",
		"float64": " %.2f ",
		"bool":    " %.2f ",
	}
)

const (
	STMT_QUAL = "qual"
)

type File struct {
	Group
	Imports map[string][]string
	Name    string
}

func NewFile(file string) *File {
	f := &File{
		Name:    file,
		Imports: map[string][]string{},
	}
	// f.Group.File = f
	return f
}

// Entidade representa menor unidade de Statement.
type Stmt struct {
	Group
	Value     interface{}
	Template  string
	Delimiter string
	// Separete string
	StmtType string
}

// Root node group of Statement.
type Group struct {
	Delimiter string
	// File  *File
	Stmts []CodeInterface
}

// Metodo realiza a renderização de um grupo e todos os seus statements.
func (g *Group) Render(buffer *bytes.Buffer) (err error) {
	var del = ""
	for _, s := range g.Stmts {
		buffer.WriteString(del)
		err = s.Render(buffer)
		del = g.Delimiter
	}
	return
}

func (f *File) GoString() string {
	return f.Group.GoString()
}

// Metodo Gera a string do arquivo.
func (g *Group) GoString() string {
	buf := bytes.Buffer{}

	if err := g.Render(&buf); err != nil {
		panic(err)
	}
	return buf.String()
}

func NewGroup() *Group {
	return &Group{}
}

func (f *File) Save() (err error) {
	var (
		out    bytes.Buffer
		stderr bytes.Buffer
	)

	if err = utils.FilePutContents(f.Name, f.GoString(), 0777); err != nil {
		return err
	}
	// fmt.Println("Slve ts")
	// Formate code with https://www.npmjs.com/package/typescript-formatter

	fmt.Println("format", f.Name)

	// cmd := exec.Command()
	// cmd.Stdin = strings.NewReader("")
	// cmd.Stdout = &out

	// cmd := exec.Command("find", "/", "-maxdepth", "1", "-exec", "wc", "-c", "{}", "\\")
	cmd := exec.Command("tsfmt", "-r", "--no-tsfmt", "--no-tslint", f.Name)
	// cmd := exec.Command("tsfmt", "-r", f.Name)
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err = cmd.Run(); err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	fmt.Println("Result: " + out.String())
	return
}

func (s *Stmt) Render(buffer *bytes.Buffer) (err error) {

	var (
		Value          interface{}
		bufLocal       *bytes.Buffer
		bufLocalMaster = &bytes.Buffer{}
		wbv            string
		delim          string
	)

	if len(s.Stmts) > 0 {

		delim = ""
		for _, stmt := range s.Stmts {
			bufLocal = &bytes.Buffer{}
			stmt.Render(bufLocal)
			bufLocalMaster.WriteString(delim + bufLocal.String())
			delim = s.Delimiter
		}
		Value = bufLocalMaster.String()
	} else {
		Value = s.Value
	}

	if strings.Contains(s.Template, "%") {

		if f, ok := Value.(func(g *Group)); ok {
			g := &Group{}
			g.Delimiter = s.Delimiter
			f(g)
			Value = g.GoString()
		}

		wbv = fmt.Sprintf(s.Template, Value)

	} else {
		wbv = s.Template
	}

	buffer.WriteString(wbv)
	return
}

func (g *Group) Lit(stmt interface{}) *Group {
	var (
		Template string
		typ      = reflect.TypeOf(stmt).Name()
	)
	if t, ok := LitTemplates[typ]; ok {
		Template = t
	} else {
		Template = " %v "
	}

	s := &Stmt{
		Template: Template,
		Value:    stmt,
	}
	g.Stmts = append(g.Stmts, s)

	return g
}

func Lit(stmt interface{}) *Group {
	return NewGroup().Lit(stmt)
}

func (g *Group) Comment(stmt string) *Group {
	var (
		Template string
	)
	if strings.Contains(stmt, "\n") {
		Template = " \n/* \n%s\n */\n "

	} else {
		Template = " \n// %s\n "
	}
	s := &Stmt{
		Template: Template,
		Value:    stmt,
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Comment(stmt string) *Group {
	return NewGroup().Comment(stmt)
}

func Dot(stmt string) *Group {
	return NewGroup().Dot(stmt)
}

func (g *Group) Dot(stmt string) *Group {

	s := &Stmt{
		Template: ".%s ",
		Value:    stmt,
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Raw(stmt string) *Group {
	return NewGroup().Raw(stmt)
}

func (g *Group) Raw(stmt string) *Group {

	s := &Stmt{
		Template: " %s ",
		Value:    stmt,
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func (g *Group) Add(gin *Group) *Group {

	s := &Stmt{
		Template: " %s ",
		Group:    *gin,
		Value:    "",
	}
	// s.Group.File = g.File
	g.Stmts = append(g.Stmts, s)
	return g
}

func (g *Group) Params(params ...CodeInterface) *Group {
	s := &Stmt{
		Template: " (%s) ",
		Group: Group{
			Stmts: params,
			// File:  g.File,
		},
		Value:     "",
		Delimiter: ", ",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func (g *Group) ParamsFun(fun func(g *Group)) *Group {
	s := &Stmt{Template: " (%s) ", Value: fun}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Params(params ...CodeInterface) *Group {
	return NewGroup().Params(params...)
}
func ParamsFun(fun func(g *Group)) *Group {
	return NewGroup().ParamsFun(fun)
}

func (g *Group) Op(op string) *Group {
	s := &Stmt{Template: "  %s  ", Value: op}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Op(op string) *Group {
	return NewGroup().Op(op)
}

func (g *Group) Line() *Group {
	s := &Stmt{
		Template: "\n",
		Value:    "",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Line() *Group {
	return NewGroup().Line()
}

func (g *Group) Block(stmts ...CodeInterface) *Group {
	s := &Stmt{
		Template: " {%s} ",
		Group: Group{
			Stmts: stmts,
		},
		Value:     "",
		Delimiter: "\n",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func (g *Group) InlineBlock(stmts ...CodeInterface) *Group {
	s := &Stmt{
		Template: " {%s} ",
		Group: Group{
			Stmts: stmts,
		},
		Value:     "",
		Delimiter: ", ",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func (g *Group) BlockFun(fun func(g *Group)) *Group {
	s := &Stmt{
		Template: "  {%s}\n ",
		Value:    fun,
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Block(stmts ...CodeInterface) *Group {
	return NewGroup().Block(stmts...)
}

func InlineBlock(stmts ...CodeInterface) *Group {
	return NewGroup().InlineBlock(stmts...)
}

func BlockFun(fun func(g *Group)) *Group {
	return NewGroup().BlockFun(fun)
}

func (g *Group) Call(params ...CodeInterface) *Group {
	s := &Stmt{
		Template: " (%s) ",
		Group: Group{
			// Stmts: []CodeInterface{params},
			Stmts: params,
		},
		Delimiter: ", ",
		Value:     "",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}
func (g *Group) CallFun(fun func(g *Group)) *Group {
	s := &Stmt{Template: " (%s) ", Value: fun}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Call(params ...CodeInterface) *Group {
	return NewGroup().Call(params...)
}
func CallFun(fun func(g *Group)) *Group {
	return NewGroup().CallFun(fun)
}

func (g *Group) Id(stmt string) *Group {
	s := &Stmt{Template: " %s ", Value: stmt}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Id(stmt string) *Group {
	return NewGroup().Id(stmt)
}

func (g *Group) Index(index ...CodeInterface) *Group {
	s := &Stmt{
		Template:  " [%s] ",
		Group:     Group{Stmts: index},
		Value:     "",
		Delimiter: ", ",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func (g *Group) IndexFun(fun func(g *Group)) *Group {
	s := &Stmt{Template: " [%s] ", Value: fun, Delimiter: ", "}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Index(index ...CodeInterface) *Group {
	return NewGroup().Index(index...)
}

func IndexFun(fun func(g *Group)) *Group {
	return NewGroup().IndexFun(fun)
}

func (g *Group) Produce(f func(g *Group)) *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    f,
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Produce(f func(g *Group)) *Group {
	return NewGroup().Produce(f)
}

type BaseCodeInterface interface {
	Render(buffer *bytes.Buffer) error
	GoString() string
	Block(stmts ...CodeInterface) *Group
	Call(params ...CodeInterface) *Group
	Comment(stmt string) *Group
	Id(stmt string) *Group
	Index(index ...CodeInterface) *Group
	Lit(stmt interface{}) *Group
	Op(op string) *Group
	Params(params ...CodeInterface) *Group
	Produce(f func(g *Group)) *Group
}

func ConvType(typ string) string {
	ntype := strings.Replace(typ, "*", "", -1)

	switch ntype {
	case "int", "int8", "int16", "int32", "int64", "float32", "float64":
		ntype = "number"
	case "bool":
		ntype = "boolean"
	case "bson.ObjectId":
		fallthrough
	case "primitive.ObjectID":
		ntype = "string"
	case "interface{}":
		ntype = "any"
	case "T":
		ntype = "T"
	default:
		if ntype[0:3] == "map" {
			parts := strings.Split(ntype, "|")
			// if parts[2] == "interface{}" {
			// 	parts[2] = "any"
			// }
			ntype = fmt.Sprintf("Map<%s,%s>", parts[1], ConvType(parts[2]))
		}
	}

	return ntype
}
