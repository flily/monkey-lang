package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/flily/monkey-lang/compiler"
	"github.com/flily/monkey-lang/evaluator"
	"github.com/flily/monkey-lang/lexer"
	"github.com/flily/monkey-lang/object"
	"github.com/flily/monkey-lang/parser"
	"github.com/flily/monkey-lang/vm"
)

var engine = flag.String("engine", "vm", "use 'vm' or 'eval'")

var input = `
	let fibonacci = fn(x) {
		if (x == 0) {
			0
		} else {
			if (x == 1) {
				return 1;
			} else {
				fibonacci(x - 1) + fibonacci(x - 2);
			}
		}
	}
	fibonacci(35);
`

func main() {
	flag.Parse()

	var duration time.Duration
	var result object.Object

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	if *engine == "vm" {
		comp := compiler.New()
		err := comp.Compile(program)
		if err != nil {
			fmt.Printf("compiler error: %s", err)
			return
		}

		machine := vm.New(comp.Bytecode())

		start := time.Now()

		err = machine.Run()
		if err != nil {
			fmt.Printf("vm error: %s", err)
		}

		duration = time.Since(start)
		result = machine.LastPoppedStackElem()

	} else {
		env := object.NewEnvironment()
		start := time.Now()
		result = evaluator.Eval(program, env)
		duration = time.Since(start)
	}

	fmt.Printf("engine=%s, result=%s, duration=%s\n",
		*engine,
		result.Inspect(),
		duration)
}
