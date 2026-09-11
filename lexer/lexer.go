package lexer

// lexer is a tiny single-purpose tokenizer for variable substitution.
//
// It mirrors the parser package's lexer style: it scans an expression string
// and emits alternating tokens of literal text and whole identifiers, so that
// substitution operates on complete identifier tokens rather than substrings.
// Function and constant names (e.g. "sin", "pi") are never touched because
// matching happens on whole tokens only.

type IdentToken struct {
	Text  string
	Ident bool // true when this token is a complete identifier
}

// lexIdentifiers scans s and returns its tokens.
func Lex(s string) []IdentToken {
	var toks []IdentToken
	i := 0
	for i < len(s) {
		if IsIdentStart(s[i]) {
			j := i + 1
			for j < len(s) && IsIdentChar(s[j]) {
				j++
			}
			toks = append(toks, IdentToken{Text: s[i:j], Ident: true})
			i = j
			continue
		}
		j := i + 1
		for j < len(s) && !IsIdentStart(s[j]) {
			j++
		}
		toks = append(toks, IdentToken{Text: s[i:j]})
		i = j
	}
	return toks
}

// firstIdent returns the first identifier token in s, or "" if none.
func FirstIdent(s string) string {
	for _, t := range Lex(s) {
		if t.Ident {
			return t.Text
		}
	}
	return ""
}

// hasIdent reports whether s contains the standalone identifier w.
func HasIdent(s, w string) bool {
	for _, t := range Lex(s) {
		if t.Ident && t.Text == w {
			return true
		}
	}
	return false
}

func IsIdentStart(ch byte) bool {
	return ch == '_' || (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func IsIdentChar(ch byte) bool {
	return IsIdentStart(ch) || (ch >= '0' && ch <= '9')
}
