package parser

// Parse lexes and parses input into an AST node without evaluating.
func Parse(input string) (Node, error) {
	l := NewLexer(input)
	toks, err := l.Tokenize()
	if err != nil {
		return nil, err
	}
	p := NewParser(toks)
	node, err := p.Parse()
	if err != nil {
		return nil, err
	}
	return node, nil
}
