package typescript

type CodeInterface interface {
	BaseCodeInterface
	TypeAliase(test string) *Group
	Any() *Group
	As() *Group
	Boolean() *Group
	Case() *Group
	Catch() *Group
	Class() *Group
	Const() *Group
	Constructor(stmt ...CodeInterface) *Group
	Debugger() *Group
	Declare() *Group
	Default() *Group
	Delete() *Group
	Do() *Group
	Else() *Group
	Enum() *Group
	Export() *Group
	Extends() *Group
	False() *Group
	Finally() *Group
	For() *Group
	From() *Group
	Function() *Group
	Get() *Group
	If(test CodeInterface) *Group
	Implements() *Group
	Import() *Group
	In() *Group
	Instanceof() *Group
	Interface() *Group
	Let() *Group
	Module() *Group
	New() *Group
	Null() *Group
	Number() *Group
	Of() *Group
	Package() *Group
	Private() *Group
	Protected() *Group
	Public() *Group
	Require() *Group
	Return() *Group
	Set() *Group
	Static() *Group
	String() *Group
	Super() *Group
	Switch() *Group
	Symbol() *Group
	This() *Group
	Throw() *Group
	True() *Group
	Try() *Group
	Type() *Group
	Typeof() *Group
	Var() *Group
	Void() *Group
	While() *Group
	With() *Group
	Yield() *Group
}

func (g *Group) TypeAliase(test string) *Group {
	s := &Stmt{Template: " <%s> ", Value: test}
	g.Stmts = append(g.Stmts, s)
	return g
}

func TypeAliase(test string) *Group {
	return NewGroup().TypeAliase(test)
}

func (g *Group) Endnl() *Group {
	s := &Stmt{Template: " ;\n ", Value: "endnl"}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Endnl() *Group {
	return NewGroup().Endnl()
}

func (g *Group) Injetable(inject string) *Group {
	s := &Stmt{Template: " @Injectable(" + inject + ")\n ", Value: "injetable"}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Injetable(inject string) *Group {
	return NewGroup().Injetable(inject)
}

func (g *Group) Any() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "any",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Any() *Group {
	return NewGroup().Any()
}

func (g *Group) As() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "as",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func As() *Group {
	return NewGroup().As()
}

func (g *Group) Boolean() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "boolean",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Boolean() *Group {
	return NewGroup().Boolean()
}

func (g *Group) Break() *Group {
	s := &Stmt{Template: " break; ", Value: "break"}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Break() *Group {
	return NewGroup().Break()
}

func (g *Group) Case() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "case",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Case() *Group {
	return NewGroup().Case()
}

func (g *Group) Catch() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "catch",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Catch() *Group {
	return NewGroup().Catch()
}

func (g *Group) Class() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "class",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Class() *Group {
	return NewGroup().Class()
}

func (g *Group) Const() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "const",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Const() *Group {
	return NewGroup().Const()
}

func (g *Group) Constructor(stmt ...CodeInterface) *Group {
	s := &Stmt{Template: " constructor(%s) ", Group: Group{Stmts: stmt}, Value: "constructor", Delimiter: ", "}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Constructor(stmt ...CodeInterface) *Group {
	return NewGroup().Constructor(stmt...)
}

func (g *Group) Continue() *Group {
	s := &Stmt{Template: " continue; ", Value: "continue"}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Continue() *Group {
	return NewGroup().Continue()
}

func (g *Group) Debugger() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "debugger",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Debugger() *Group {
	return NewGroup().Debugger()
}

func (g *Group) Declare() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "declare",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Declare() *Group {
	return NewGroup().Declare()
}

func (g *Group) Default() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "default",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Default() *Group {
	return NewGroup().Default()
}

func (g *Group) Delete() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "delete",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Delete() *Group {
	return NewGroup().Delete()
}

func (g *Group) Do() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "do",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Do() *Group {
	return NewGroup().Do()
}

func (g *Group) Else() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "else",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Else() *Group {
	return NewGroup().Else()
}

func (g *Group) Endl() *Group {
	s := &Stmt{Template: " ; ", Value: "endl"}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Endl() *Group {
	return NewGroup().Endl()
}

func (g *Group) Enum() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "enum",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Enum() *Group {
	return NewGroup().Enum()
}

func (g *Group) Export() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "export",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Export() *Group {
	return NewGroup().Export()
}

func (g *Group) Extends() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "extends",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Extends() *Group {
	return NewGroup().Extends()
}

func (g *Group) False() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "false",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func False() *Group {
	return NewGroup().False()
}

func (g *Group) Finally() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "finally",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Finally() *Group {
	return NewGroup().Finally()
}

func (g *Group) For() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "for",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func For() *Group {
	return NewGroup().For()
}

func (g *Group) From() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "from",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func From() *Group {
	return NewGroup().From()
}

func (g *Group) Function() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "function",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Function() *Group {
	return NewGroup().Function()
}

func (g *Group) Get() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "get",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Get() *Group {
	return NewGroup().Get()
}

func (g *Group) If(test CodeInterface) *Group {
	s := &Stmt{Template: " if (%s) ", Group: Group{Stmts: []CodeInterface{test}}, Value: "if"}
	g.Stmts = append(g.Stmts, s)
	return g
}

func If(test CodeInterface) *Group {
	return NewGroup().If(test)
}

func (g *Group) Implements() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "implements",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Implements() *Group {
	return NewGroup().Implements()
}

func (g *Group) Import() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "import",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Import() *Group {
	return NewGroup().Import()
}

func (g *Group) In() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "in",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func In() *Group {
	return NewGroup().In()
}

func (g *Group) Instanceof() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "instanceof",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Instanceof() *Group {
	return NewGroup().Instanceof()
}

func (g *Group) Interface() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "interface",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Interface() *Group {
	return NewGroup().Interface()
}

func (g *Group) Let() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "let",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Let() *Group {
	return NewGroup().Let()
}

func (g *Group) Module() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "module",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Module() *Group {
	return NewGroup().Module()
}

func (g *Group) New() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "new",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func New() *Group {
	return NewGroup().New()
}

func (g *Group) Null() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "null",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Null() *Group {
	return NewGroup().Null()
}

func (g *Group) Number() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "number",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Number() *Group {
	return NewGroup().Number()
}

func (g *Group) Of() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "of",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Of() *Group {
	return NewGroup().Of()
}

func (g *Group) Package() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "package",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Package() *Group {
	return NewGroup().Package()
}

func (g *Group) Private() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "private",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Private() *Group {
	return NewGroup().Private()
}

func (g *Group) Protected() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "protected",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Protected() *Group {
	return NewGroup().Protected()
}

func (g *Group) Public() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "public",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Public() *Group {
	return NewGroup().Public()
}

func (g *Group) Require() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "require",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Require() *Group {
	return NewGroup().Require()
}

func (g *Group) Return() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "return",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Return() *Group {
	return NewGroup().Return()
}

func (g *Group) Set() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "set",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Set() *Group {
	return NewGroup().Set()
}

func (g *Group) Static() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "static",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Static() *Group {
	return NewGroup().Static()
}

func (g *Group) String() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "string",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func String() *Group {
	return NewGroup().String()
}

func (g *Group) Super() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "super",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Super() *Group {
	return NewGroup().Super()
}

func (g *Group) Switch() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "switch",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Switch() *Group {
	return NewGroup().Switch()
}

func (g *Group) Symbol() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "symbol",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Symbol() *Group {
	return NewGroup().Symbol()
}

func (g *Group) This() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "this",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func This() *Group {
	return NewGroup().This()
}

func (g *Group) Throw() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "throw",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Throw() *Group {
	return NewGroup().Throw()
}

func (g *Group) True() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "true",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func True() *Group {
	return NewGroup().True()
}

func (g *Group) Try() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "try",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Try() *Group {
	return NewGroup().Try()
}

func (g *Group) Type() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "type",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Type() *Group {
	return NewGroup().Type()
}

func (g *Group) Typeof() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "typeof",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Typeof() *Group {
	return NewGroup().Typeof()
}

func (g *Group) Var() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "var",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Var() *Group {
	return NewGroup().Var()
}

func (g *Group) Void() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "void",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Void() *Group {
	return NewGroup().Void()
}

func (g *Group) While() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "while",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func While() *Group {
	return NewGroup().While()
}

func (g *Group) With() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "with",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func With() *Group {
	return NewGroup().With()
}

func (g *Group) Yield() *Group {
	s := &Stmt{
		Template: " %s ",
		Value:    "yield",
	}
	g.Stmts = append(g.Stmts, s)
	return g
}

func Yield() *Group {
	return NewGroup().Yield()
}
