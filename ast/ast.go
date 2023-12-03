// Package ast is ast class
package ast

import "github.com/aaazlkm/go-interpreter/token"

type Node interface {
	// ノードに紐づくトークンのリテラル値を返す
	TokenLiteral() string
}

type Statement interface {
	Node

	// 空のメソッド
	// マーカーインターフェースと呼ばれ、特定の型がステートメントであることを示す。
	statementNode()
}

type Expression interface {
	Node

	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		const x = 10
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	LetToken token.Token
	Name     string
	Value    Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string { return ls.Value.TokenLiteral() }

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) statementNode() {}

func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
