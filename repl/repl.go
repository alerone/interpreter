package repl

import (
	"bufio"
	"fmt"
	"interpreter/evaluator"
	"interpreter/lexer"
	"interpreter/object"
	"interpreter/parser"
	"io"
)

const PROMPT = "8=> "
const MONKEY_FACE = `       _
      (_)          _
  _         .=.   (_)
 (_)   _   //(')_
      //'\/ |\ 0'\\
      ||-.\_|_/.-||
      )/ |_____| \(    _
     0   #/\ /\#  0   (_)
        _| o o |_
 _     ((|, ^ ,|))
(_)     '||\_/||'
         || _ ||      _
         | \_/ |     (_)
     0.__.\   /.__.0
      '._  '"'  _.'
         / ;  \ \
       0'-' )/''-0
           0'
`
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
    env := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! we ran into some monkey business here!")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
