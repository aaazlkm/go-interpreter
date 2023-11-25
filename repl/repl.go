package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/aaazlkm/go-interpreter/lexer"
	"github.com/aaazlkm/go-interpreter/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf("%s\n", PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			break
		}

		line := scanner.Text()
		l := lexer.New(line)

		// そうか、これシンプルなfor文なのか。
		// だから、初期化; 条件式; 後処理式 という形式になっている。
		// 条件式のように勝手に読んでしまった。
		for tk := l.NextToken(); tk.Type != token.EOF; tk = l.NextToken() {
			fmt.Printf("%+v\n", tk)
		}
	}
}
