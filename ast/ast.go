// Package ast defines the Abstract Syntax Tree structures for the language

package ast

// Node is the base interface for all AST nodes
// Every node in the AST must implement TokenLiteral()
import (
	"bytes"
	"lang/token"
	"strings"
)

type Node interface {
	TokenLiteral() string
	String() string
}

// Statement represents nodes that don't produce values
// Examples: let statements, return statements
type Statement interface {
	Node
	statementNode() // Marker method to distinguish statements
}

// Expression represents nodes that produce values
// Examples: identifiers, literals, function calls
type Expression interface {
	Node
	expressionNode() // Marker method to distinguish expressions
}

// Program is the root node of every AST
// Contains a slice of statements that make up the program
type Program struct {
	Statements []Statement
}

type Boolean struct {
	Token token.Token
	Value bool
}

// TokenLiteral returns the first statement's literal value
// Returns empty string for empty programs
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LetStatement represents a variable declaration
// Example: let x = 5;
type LetStatement struct {
	Token token.Token // The 'let' token
	Name  *Identifier // The variable name
	Value Expression  // The value being assigned
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

type PrefixExpression struct {
	Token    token.Token // The prefix token, e.g. !
	Operator string
	Right    Expression
}

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (pe *PrefixExpression) expressionNode() { /*these are the markers to distinguish between statement and expression*/
}

func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}
func (es *ExpressionStatement) statementNode()       { /* .. */ }
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (ls *LetStatement) statementNode() { /*these are the markers to distinguish between statement and expression*/
}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier represents a name in the program
// Example: variable names, function names
type Identifier struct {
	Token token.Token // The identifier token
	Value string      // The actual name
}

func (i *Identifier) expressionNode() { /*these are the markers to distinguish between statement and expression*/
}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// Key concepts:
// 1. Every node implements Node interface
// 2. Two main types of nodes:
//    - Statements (don't produce values)
//    - Expressions (produce values)
// 3. Program is the root node containing all statements
// 4. Each node holds its token for error reporting and debugging
// 5. Marker methods (statementNode/expressionNode) help with type safety

type ReturnStatement struct {
	Token       token.Token // The 'return' token
	ReturnValue Expression  // The value being returned
}

func (rs *ReturnStatement) statementNode()       { /*...*/ }
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

func (i *Identifier) String() string { return i.Value }

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      { /*...*/ }
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (oe *InfixExpression) expressionNode() { /*...*/ }

func (oe *InfixExpression) TokenLiteral() string { return oe.Token.Literal }

func (oe *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" " + oe.Operator + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")

	return out.String()
}

func (b *Boolean) expressionNode()      { /*...*/ }
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }

func (ie *IfExpression) expressionNode() { /*...*/ }

func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }

func (ie *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

func (bs *BlockStatement) statementNode() { /* --- */ }

func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode() { /* --- */ }

func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }

func (f1 *FunctionLiteral) String() string {
	var out bytes.Buffer
	params := []string{}

	for _, p := range f1.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(f1.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	out.WriteString(f1.Body.String())

	return out.String()
}

type CallExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func (ce *CallExpression) expressionNode()      { /* --- */ }
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }

func (ce *CallExpression) String() string {
	var out bytes.Buffer
	args := []string{}

	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}
