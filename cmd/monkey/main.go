package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/user"
	"strings"

	"github.com/flily/monkey-lang/compiler"
	"github.com/flily/monkey-lang/evaluator"
	"github.com/flily/monkey-lang/lexer"
	"github.com/flily/monkey-lang/object"
	"github.com/flily/monkey-lang/parser"
	"github.com/flily/monkey-lang/vm"
)

const (
	PROMPT      = ">> "
	MONKEY_FACE = `
           __,__
  .--.  .-"     "-.  .--.
 / .. \/  .-. .-.  \/ .. \
| |  '|  /   Y   \  |'  | |
| \   \  \ 0 | 0 /  /   / |
 \ '- ,\.-"""""""-./, -' /
  ''-' /_   ^ ^   _\ '-''
      |  \._   _./  |
      \   \ '~' /   /
       '._ '-=-' _.'
          '-----'
`
)

func getBanner() string {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}

	lines := []string{
		fmt.Sprintf("Hello %s! This is the Monkey programming language!\n", u.Username),
	}

	return strings.Join(lines, "")
}

func runWithInterpreter(code string, env *object.Environment) (object.Object, []string) {
	l := lexer.New(code)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		return nil, p.Errors()
	}

	result := evaluator.Eval(program, env)
	return result, nil
}

type compilerContext struct {
	constants   []object.Object
	globals     []object.Object
	symbolTable *compiler.SymbolTable
}

func (c *compilerContext) initBuiltins() {
	for i, v := range object.Builtins {
		c.symbolTable.DefineBuiltin(i, v.Name)
	}
}

func newCompilerContext() *compilerContext {
	c := &compilerContext{
		constants:   []object.Object{},
		globals:     make([]object.Object, vm.GlobalsSize),
		symbolTable: compiler.NewSymbolTable(),
	}

	c.initBuiltins()
	return c
}

func runWithCompiler(code string, ctx *compilerContext) (object.Object, []string) {
	l := lexer.New(code)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		return nil, p.Errors()
	}

	comp := compiler.NewWithState(ctx.symbolTable, ctx.constants)
	err := comp.Compile(program)
	if err != nil {
		return nil, []string{err.Error()}
	}

	bytecode := comp.Bytecode()
	ctx.constants = bytecode.Constants

	machine := vm.NewWithGlobalsStore(bytecode, ctx.globals)
	err = machine.Run()
	if err != nil {
		return nil, []string{err.Error()}
	}

	result := machine.LastPoppedStackElem()
	return result, nil
}

type MonkeyConfigure struct {
	In     io.Reader
	Out    io.Writer
	Prompt string

	Interactive bool
}

func NewMonkeyConfigure() *MonkeyConfigure {
	c := &MonkeyConfigure{
		In:     os.Stdin,
		Out:    os.Stdout,
		Prompt: PROMPT,
	}

	return c
}

func (c *MonkeyConfigure) StartEvaluator(files []string) {
	scanner := NewInteractiveScanner(c.In, c.Out, c.Prompt)
	scanner.AddFiles(files)

	env := object.NewEnvironment()

	for {
		line, ok, err := scanner.Scan()
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
			break
		}

		if !ok {
			break
		}

		result, errMessages := runWithInterpreter(line, env)
		if len(errMessages) > 0 {
			scanner.WriteLines([]string{
				MONKEY_FACE,
				"Woops! We ran into some monkey business here!",
				" parser errors:",
			})
			scanner.WriteLines(errMessages)
			continue
		}

		if result != nil {
			scanner.WriteString(result.Inspect())
			scanner.WriteString("\n")
		}
	}

	fmt.Println("exit.")
}

func (c *MonkeyConfigure) StartCompiler(files []string) {
	scanner := NewInteractiveScanner(c.In, c.Out, c.Prompt)
	scanner.AddFiles(files)

	ctx := newCompilerContext()

	for {
		line, ok, err := scanner.Scan()
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
			break
		}

		if !ok {
			break
		}

		result, errMessages := runWithCompiler(line, ctx)
		if len(errMessages) > 0 {
			scanner.WriteLines([]string{
				MONKEY_FACE,
				"Woops! We ran into some monkey business here!",
				" parser errors:",
			})
			scanner.WriteLines(errMessages)
			continue
		}

		if result != nil {
			scanner.WriteString(result.Inspect())
			scanner.WriteString("\n")
		}
	}
}

func main() {
	modInteractive := flag.Bool("i", false, "run as interactive mode, enabled when no file is provided")
	engine := flag.String("e", "vm", "engine to execute monkey scripts, `vm` or `eval`, default is `vm`")

	flag.Parse()
	args := flag.Args()

	conf := NewMonkeyConfigure()

	conf.Interactive = *modInteractive

	switch *engine {
	case "vm":
		conf.StartEvaluator(args)

	case "eval":
		conf.StartCompiler(args)

	default:
		fmt.Printf("unknown engine: %s\n", *engine)
	}
}
