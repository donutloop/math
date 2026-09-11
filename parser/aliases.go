package parser

import "prototype_kl/lexer"

// Type aliases for symbols that moved to the shared lexer package.
type (
	Token      = lexer.Token
	Lexer      = lexer.Lexer
	ParseError = lexer.ParseError
	EvalError  = lexer.EvalError
)

// NewLexer constructs an expression tokenizer.
var NewLexer = lexer.NewLexer

// Op constants.
const (
	OpAdd        = lexer.OpAdd
	OpSub        = lexer.OpSub
	OpMul        = lexer.OpMul
	OpDiv        = lexer.OpDiv
	OpMod        = lexer.OpMod
	OpPower      = lexer.OpPower
	OpEQ         = lexer.OpEQ
	OpNE         = lexer.OpNE
	OpLT         = lexer.OpLT
	OpLE         = lexer.OpLE
	OpGT         = lexer.OpGT
	OpGE         = lexer.OpGE
	OpAND        = lexer.OpAND
	OpOR         = lexer.OpOR
	OpANDAND     = lexer.OpANDAND
	OpOROR       = lexer.OpOROR
	OpShiftLeft  = lexer.OpShiftLeft
	OpShiftRight = lexer.OpShiftRight
	OpTilde      = lexer.OpTilde
	OpPercent    = lexer.OpPercent
	OpLParen     = lexer.OpLParen
	OpRParen     = lexer.OpRParen
	OpComma      = lexer.OpComma
	OpColon      = lexer.OpColon
)

// TokenType constants.
const (
	TokenAND        = lexer.TokenAND
	TokenBitAND     = lexer.TokenBitAND
	TokenBitOR      = lexer.TokenBitOR
	TokenColon      = lexer.TokenColon
	TokenComma      = lexer.TokenComma
	TokenConstant   = lexer.TokenConstant
	TokenDivide     = lexer.TokenDivide
	TokenEOF        = lexer.TokenEOF
	TokenEQ         = lexer.TokenEQ
	TokenFactorial  = lexer.TokenFactorial
	TokenFunction   = lexer.TokenFunction
	TokenGE         = lexer.TokenGE
	TokenGT         = lexer.TokenGT
	TokenLE         = lexer.TokenLE
	TokenLParen     = lexer.TokenLParen
	TokenLT         = lexer.TokenLT
	TokenMinus      = lexer.TokenMinus
	TokenModulo     = lexer.TokenModulo
	TokenMultiply   = lexer.TokenMultiply
	TokenNE         = lexer.TokenNE
	TokenNumber     = lexer.TokenNumber
	TokenOR         = lexer.TokenOR
	TokenPercent    = lexer.TokenPercent
	TokenPlus       = lexer.TokenPlus
	TokenPower      = lexer.TokenPower
	TokenQuestion   = lexer.TokenQuestion
	TokenRParen     = lexer.TokenRParen
	TokenShiftLeft  = lexer.TokenShiftLeft
	TokenShiftRight = lexer.TokenShiftRight
	TokenTilde      = lexer.TokenTilde
)

// TokenType constants.
// Shared maps and errors.
var (
	SupportedFunctions  = lexer.SupportedFunctions
	SupportedConstants  = lexer.SupportedConstants
	DefaultPrecision    = lexer.DefaultPrecision
	ErrUnknownFunction  = lexer.ErrUnknownFunction
	ErrInvalidCharacter = lexer.ErrInvalidCharacter
	ErrBadArity         = lexer.ErrBadArity
	ErrDivisionByZero   = lexer.ErrDivisionByZero
	ErrSqrtNegative     = lexer.ErrSqrtNegative
	ErrOverflow         = lexer.ErrOverflow
	ErrFactorial        = lexer.ErrFactorial
	ErrDomain           = lexer.ErrDomain
	ErrEmptyExpression  = lexer.ErrEmptyExpression
	ErrMismatchedParen  = lexer.ErrMismatchedParen
	ErrUnexpectedToken  = lexer.ErrUnexpectedToken
)
