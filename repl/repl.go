package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/watsoncj/monkey/lexer"
	"github.com/watsoncj/monkey/parser"
	"github.com/watsoncj/monkey/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}

		p := parser.New(l)
		program := p.ParseProgram()
		fmt.Fprintf(out, "Parsed (%d statements):\n%+v\n", len(program.Statements), program.String())
	}
}
