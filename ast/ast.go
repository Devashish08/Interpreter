// Package ast defines the Abstract Syntax Tree structures for the language

package ast

// Node is the base interface for all AST nodes
// Every node in the AST must implement TokenLiteral()
import (
	"lang/token"
)

type Node interface {
	TokenLiteral() string
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
