package calc

import (
	"bufio"
	"io"
)

// state.go — the Calculator runtime state: variables, previous answer, and
// user-defined functions. Grouped here so the evaluation and construct
// machinery live in focused files by language-design concern.

type Calculator struct {
	vars        map[string]float64
	funcs       map[string]*funcDef
	ans         float64
	hasAns      bool
	memory      float64
	hasMem      bool
	history     []string
	results     []string
	statePath   string
	degMode     bool
	gradMode    bool
	sci         bool
	prec        int
	undoStack   []state
	redoStack   []state
	quiet       bool
	quietAssign bool
	jsonMode    bool
	jsonlMode   bool
	csvMode     bool
	lastExpr    string
	eng         bool
	expandDepth int
	errCount    int
	base        int
	in          *bufio.Reader
	out         io.Writer
}
