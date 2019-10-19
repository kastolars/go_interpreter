package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/kstola2/monkey_interpreter/lexer"
	"github.com/kstola2/monkey_interpreter/mytoken"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != mytoken.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}

}
