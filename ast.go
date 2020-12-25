package main

// abstract syntax tree for statements
type Ast struct {
    Statements []*Statement
}

type AstKind uint

// kinds of instructions
const (
	SelectKind AstKind = iota
	CreateTableKind
	InsertKind
)

// statetements
type Statement struct {
    SelectStatement      *SelectStatement
    CreateTableStatement *CreateTableStatement
    InsertStatement      *InsertStatement
    Kind                 AstKind
}

// INSERT
type InsertStatement struct {
    table  token
    values *[]*expression
}

type expressionKind uint

const (
    literalKind expressionKind = iota
)

type expression struct {
    literal *token
    kind    expressionKind
}

// CREATE
type columnDefinition struct {
    name     token
    datatype token
}

type CreateTableStatement struct {
    name token
    cols *[]*columnDefinition
}

type SelectStatement struct {
    item []*expression
    from token
}