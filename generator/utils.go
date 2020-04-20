package generator

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

type File struct {
	Group
	Name string
}

func NewFile(file string) *File {
	return &File{Name: file}
}

// Entidade representa menor unidade de Statement.
type Stmt struct {
	Group
	Value    interface{}
	Template string
	Separete string
}

// Root node group of Statement.
type Group struct {
	Stmts []CodeInterface
}

// Metodo realiza a renderização de um grupo e todos os seus statements.
func (g *Group) Render(buffer *bytes.Buffer) (err error) {
	for _, s := range g.Stmts {
		err = s.Render(buffer)
	}
	return
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

	fmt.Println("Save in :", f.Name)

	if err = utils.FilePutContents(f.Name, f.GoString(), 777); err != nil {
		return
	}
	// Formate code with https://www.npmjs.com/package/typescript-formatter

	var out bytes.Buffer

	cmd := exec.Command("tsfmt", "-r", f.Name)
	cmd.Stdin = strings.NewReader("")
	cmd.Stdout = &out

	if err = cmd.Run(); err != nil {
		panic(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
	return
}

func (s *Stmt) Render(buffer *bytes.Buffer) (err error) {
	// fmt.Println("Render stmt")
	// spew.Dump(s)

	var (
		Value          interface{}
		bufLocal       *bytes.Buffer
		bufLocalMaster = &bytes.Buffer{}
		wbv            string
	)

	if len(s.Stmts) > 0 {
		for _, stmt := range s.Stmts {
			bufLocal = &bytes.Buffer{}
			stmt.Render(bufLocal)
			bufLocalMaster.WriteString(bufLocal.String())
		}
		Value = bufLocalMaster.String()
	} else {
		Value = s.Value
	}

	if strings.Contains(s.Template, "%") {
		// switch Value.(type) {
		// case string:
		// 	wbv = fmt.Sprintf(s.Template, Value.(string))
		// case :
		// }
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
		Template = " \n/*\n%s\n*/\n "

	} else {
		Template = " \n//%s\n "
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

func (g *Group) Params(params ...CodeInterface) *Group {
	s := &Stmt{Template: " (%s) ", Group: Group{Stmts: params}, Value: "params"}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Params(params ...CodeInterface) *Group {
	return NewGroup().Params(params...)
}

func (g *Group) Op(op string) *Group {
	s := &Stmt{Template: "  %s  ", Value: op}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Op(op string) *Group {
	return NewGroup().Op(op)
}

func (g *Group) Block(stmts ...CodeInterface) *Group {
	s := &Stmt{Template: "  {%s}\n ", Group: Group{Stmts: stmts}, Value: "block"}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Block(stmts ...CodeInterface) *Group {
	return NewGroup().Block(stmts...)
}

func (g *Group) Call(params CodeInterface) *Group {
	s := &Stmt{Template: " (%s) ", Group: Group{Stmts: []CodeInterface{params}}, Value: "call"}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Call(params CodeInterface) *Group {
	return NewGroup().Call(params)
}

func (g *Group) Id(stmt string) *Group {
	s := &Stmt{Template: " %s ", Value: stmt}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Id(stmt string) *Group {
	return NewGroup().Id(stmt)
}

func (g *Group) Index(index CodeInterface) *Group {
	s := &Stmt{Template: " [%s] ", Group: Group{Stmts: []CodeInterface{index}}, Value: "index"}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Index(index CodeInterface) *Group {
	return NewGroup().Index(index)
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

type CodeInterface interface {
	Render(buffer *bytes.Buffer) error
	GoString() string
	Block(stmts ...CodeInterface) *Group
	Call(params CodeInterface) *Group
	Comment(stmt string) *Group
	Id(stmt string) *Group
	Index(index CodeInterface) *Group
	Lit(stmt interface{}) *Group
	Op(op string) *Group
	Params(params ...CodeInterface) *Group
	Produce(f func(g *Group)) *Group
}
